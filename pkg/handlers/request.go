package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/glitch"
	"github.com/gorilla/mux"
	"github.com/music-tribe/uuid"
)

// Request - methods required to satisfy interface
type Request interface {
	BindRequestBody() Request
	BindQueryParams() Request
	ReturnError() error
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
		r.Err = glitch.NewError(r.Err, 400)
		return r
	}
	URLParams := mux.Vars(r.HTTPRequest)
	JSONmarshallingBuffer := bytes.NewBuffer([]byte{})
	r.Err = json.NewEncoder(JSONmarshallingBuffer).Encode(URLParams)
	if r.Err == nil {
		r.Err = glitch.NewError(r.Err, 400)
		return r
	}
	r.Err = json.NewDecoder(JSONmarshallingBuffer).Decode(&r)
	return r
}

// BindQueryParams -
func (r *request) BindQueryParams() Request {
	if r.Err != nil {
		r.Err = glitch.NewError(r.Err, http.StatusBadRequest)
		return r
	}
	// bind the query params - remove values from arrays
	r.QueryParams = removeMapValuesFromArrays(r.HTTPRequest.URL.Query())
	return r
}
func (r *request) BindRequestBody() Request {
	if r.Err != nil {
		return r
	}
	// bind request body - depending on type ie. multipart form or JSON
	return r
}
