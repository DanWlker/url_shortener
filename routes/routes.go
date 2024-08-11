package routes

import (
	"log"
	"net/http"
)

func RegisterRoutes(
	mux *http.ServeMux,
	logger *log.Logger,
) error {
	// pages
	mux.HandleFunc("/", NewHomeHandler(logger))

	// functions
	mux.HandleFunc("POST /shorten", ShortenHandler(logger))

	return nil
}
