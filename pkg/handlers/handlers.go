package handlers

import (
	"fmt"
	"net/http"

	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/database"
	"github.com/music-tribe/uuid"
)

// Handlers holds a representation of the required handler methods
type Handlers interface {
	UploadImage(http.ResponseWriter, *http.Request) error
}

// Request requires methods for it to be satisfied
type Request interface {
	BindRequestBody() Request
}

// handlers will satisfy the Handlers interface
type handlers struct {
	database.Database
}

type request struct {
	HTTPRequest *http.Request
	FileUUID    uuid.UUID
	Filename    string
	FileData    []byte
}

func newRequest(r *http.Request) *request {
	return &request{
		HTTPRequest: r,
	}
}
func (r *request) GetFileID() uuid.UUID { return r.FileUUID }
func (r *request) GetFilename() string  { return r.Filename }
func (r *request) GetFileData() []byte  { return r.FileData }
func (r *request) BindURLParams() Request {
	// bind request params
	return r
}
func (r *request) BindQueryParams() Request {
	// bind any query params
	return r
}
func (r *request) BindRequestBody() Request {
	// bind request body
	//// depending on type ie. multipart form or JSON
	return r
}

// UploadImage takes a multipart image file and stores it in the db
// the image will also be displayed to the user
func (h *handlers) UploadImage(w http.ResponseWriter, r *http.Request) error {
	uploadRequest := newRequest(r)
	fmt.Printf("%v+", uploadRequest)
	return nil
}
