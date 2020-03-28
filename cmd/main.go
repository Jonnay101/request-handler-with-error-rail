package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/database"
	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/handlers"
)

func main() {
	// get env vars
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// create required utilities
	imageRepository := database.CreateNewImageRepository("tmp/images")
	handlers := handlers.NewHandlers(imageRepository)
	// set routes
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/{person}", hello)
	muxRouter.HandleFunc("/images/user/{user_uuid}/image/{image_uuid}", handlers.UploadImage)
	// create and start server
	srv := &http.Server{
		Handler:      muxRouter,
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// hello test handler
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, this is the thingy!!!")
}
