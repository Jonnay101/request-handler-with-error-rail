package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/database"
	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/handlers"
	"github.com/Jonnay101/request-handler-with-error-rail.git/pkg/routes"
)

func main() {
	// get relevant variables for the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// create required utilities
	imageRepository := database.CreateNewImageRepository("tmp/images")
	handlers := handlers.NewHandlers(imageRepository)
	// set up router with attached routes
	muxRouter := routes.NewRoutes(handlers)
	// create and start server
	srv := &http.Server{
		Handler:      muxRouter,
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
