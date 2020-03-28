package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/database"
	"github.com/gorilla/mux"
	"github.com/music-tribe/uuid"
)

// Handlers - methods required to satisfy interface
type Handlers interface {
	UploadImage(http.ResponseWriter, *http.Request)
}

// Request - methods required to satisfy interface
type Request interface {
	BindRequestBody() Request
	BindQueryParams() Request
	ReturnError() error
}

// handlers will satisfy the Handlers interface
type handlers struct {
	database.Database
}

// NewHandlers returns a set of methods for handling image requests
func NewHandlers(db database.Database) Handlers {
	return &handlers{}
}

type request struct {
	HTTPRequest *http.Request     `json:"http_request,omitempty"`
	UserUUID    uuid.UUID         `json:"user_uuid,omitempty" bson:"user_uuid"`
	FileUUID    uuid.UUID         `json:"file_uuid,omitempty" bson:"file_uuid"`
	Filename    string            `json:"filename,omitempty"`
	FileData    []byte            `json:"file_data,omitempty"`
	QueryParams map[string]string `json:"query_params,omitempty"`
	Err         error             `json:"-"`
}

// newRequest creates a new request object
func newRequest(r *http.Request) *request { return &request{HTTPRequest: r} }

// request interaction functions
func (r *request) GetFileID() uuid.UUID { return r.FileUUID }
func (r *request) GetFilename() string  { return r.Filename }
func (r *request) GetFileData() []byte  { return r.FileData }
func (r *request) ReturnError() error   { return r.Err }

// BindURLParams encodes the map of url params into JSON
// it then decodes this JSON into our request struct
func (r *request) BindURLParams() Request {
	if r.Err != nil {
		return r
	}
	URLParams := mux.Vars(r.HTTPRequest)
	JSONmarshallingBuffer := bytes.NewBuffer([]byte{})
	r.Err = json.NewEncoder(JSONmarshallingBuffer).Encode(URLParams)
	if r.Err == nil {
		json.NewDecoder(JSONmarshallingBuffer).Decode(&r)
	}
	return r
}

// BindQueryParams -
func (r *request) BindQueryParams() Request {
	if r.Err == nil {
		// bind the query params - remove values from arrays
		r.QueryParams = removeMapValuesFromArrays(r.HTTPRequest.URL.Query())
	}
	return r
}
func (r *request) BindRequestBody() Request {
	if r.Err != nil {
		return r
	}
	// bind request body - depending on type ie. multipart form or JSON
	return r
}

// UploadImage takes a multipart image file and stores it in the db
// the image will also be displayed to the user
func (h *handlers) UploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is hit")
	uploadRequest := newRequest(r).
		BindURLParams().
		BindQueryParams()
	fmt.Fprintf(w, "%v", uploadRequest)
}

func removeMapValuesFromArrays(initialMap map[string][]string) map[string]string {
	flattenedMap := make(map[string]string)
	for key, value := range initialMap {
		flattenedMap[key] = value[0]
	}
	return flattenedMap
}
