package routes

import (
	"log"
	"net/http"

	"url_shortener/storage"
)

func RegisterRoutes(
	mux *http.ServeMux,
	logger *log.Logger,
	storageClient storage.StorageClient,
) error {
	// pages
	mux.HandleFunc("/", NewHomeHandler(logger, storageClient))

	// functions
	mux.HandleFunc("POST /shorten", ShortenHandler(logger, storageClient))

	return nil
}
