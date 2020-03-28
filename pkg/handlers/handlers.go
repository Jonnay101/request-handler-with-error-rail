package handlers

import (
	"fmt"
	"net/http"

	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/database"
	"github.com/music-tribe/uuid"
)

// Handlers - methods required to satisfy interface
type Handlers interface {
	UploadImage(http.ResponseWriter, *http.Request) error
}

// Request - methods required to satisfy interface
type Request interface {
	BindRequestBody() Request
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
	UserUUID    uuid.UUID     `json:"user_uuid,omitempty"`
	FileUUID    uuid.UUID     `json:"file_uuid,omitempty"`
	Filename    string        `json:"filename,omitempty"`
	FileData    []byte        `json:"file_data,omitempty"`
}

func newRequest(r *http.Request) *request { return &request{HTTPRequest: r} }
func (r *request) GetFileID() uuid.UUID   { return r.FileUUID }
func (r *request) GetFilename() string    { return r.Filename }
func (r *request) GetFileData() []byte    { return r.FileData }
func (r *request) BindURLParams() Request {
	// bind request params
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
func (h *handlers) UploadImage(w http.ResponseWriter, r *http.Request) error {
	uploadRequest := newRequest(r)
	fmt.Printf("%v+", uploadRequest)
	return nil
}
