package handlers

import (
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
	HTTPRequest *http.Request `json:"http_request,omitempty"`
	UserUUID    uuid.UUID     `json:"user_uuid,omitempty" bson:"user_uuid"`
	FileUUID    uuid.UUID     `json:"file_uuid,omitempty" bson:"file_uuid"`
	Filename    string        `json:"filename,omitempty"`
	FileData    []byte        `json:"file_data,omitempty"`
	Err         error         `json:"_, omitempty"`
}

func newRequest(r *http.Request) *request { return &request{HTTPRequest: r} }
func (r *request) GetFileID() uuid.UUID   { return r.FileUUID }
func (r *request) GetFilename() string    { return r.Filename }
func (r *request) GetFileData() []byte    { return r.FileData }
func (r *request) ReturnError() error     { return r.Err }
func (r *request) BindURLParams() Request {
	// bind request params
	mapOfDynaimcParams := mux.Vars(r.HTTPRequest)
	var byt []byte
	// marshal map of params to json
	byt, r.Err = json.Marshal(mapOfDynaimcParams)
	if r.Err != nil {
		return r
	}
	// unmarshal the json to our request struct
	r.Err = json.Unmarshal(byt, &r)
	return r
}
func (r *request) BindQueryParams() Request {
	// bind any query params
	return r
}
func (r *request) BindRequestBody() Request {
	// bind request body - depending on type ie. multipart form or JSON
	return r
}

// UploadImage takes a multipart image file and stores it in the db
// the image will also be displayed to the user
func (h *handlers) UploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is hit")
	uploadRequest := newRequest(r).BindURLParams()
	fmt.Fprintf(w, "%v", uploadRequest)
}
