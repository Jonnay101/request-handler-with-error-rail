package handlers

import (
	"fmt"
	"net/http"

	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/database"
)

// Handlers - methods required to satisfy interface
type Handlers interface {
	UploadImage(http.ResponseWriter, *http.Request)
}

// handlers will satisfy the Handlers interface
type handlers struct {
	database.Database
}

// NewHandlers returns a set of methods for handling image requests
func NewHandlers(db database.Database) Handlers {
	return &handlers{}
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
