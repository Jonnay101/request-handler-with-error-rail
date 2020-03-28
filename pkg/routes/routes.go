package routes

import (
	"fmt"
	"net/http"

	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/handlers"
	"github.com/gorilla/mux"
)

// NewRoutes returns a group of routes
func NewRoutes(handlers handlers.Handlers) (mr *mux.Router) {
	mr = mux.NewRouter()
	// set routes
	mr.HandleFunc("/{person}", hello).Methods("GET")
	mr.HandleFunc("/images/user/{user_uuid}/image/{image_uuid}", handlers.UploadImage).Methods("POST")
	return
}

// hello test handler
func hello(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := "someone"
	if person, ok := params["person"]; ok {
		name = person
	}

	fmt.Fprintf(w, "Hello %s!!!", name)
}
