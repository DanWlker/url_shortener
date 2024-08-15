package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/DanWlker/url_shortener/storage"
)

func return404(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusNotFound)
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return fmt.Errorf("template.ParseFiles: %w", err)
	}
	if err := t.Execute(w, nil); err != nil {
		return fmt.Errorf("t.Execute: %w", err)
	}
	return nil
}

func NewHomeHandler(logger *log.Logger, storageClient storage.StorageClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			tmpl, err := template.ParseFiles("templates/index.html")
			if err != nil {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				logger.Printf("Path: / template.ParseFiles: %v\n", err)
				return
			}
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				logger.Printf("Path: / tmpl.Execute: %v\n", err)
				return
			}
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 2 {
			if err := return404(w); err != nil {
				logger.Printf("Path: / return404: %v\n", err)
			}
			return
		}

		id64, err := strconv.ParseInt(parts[1], 16, 64)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			logger.Printf("Path: / strconv.ParseInt: %v\n", err)
			return
		}

		url, err := storageClient.Retrieve(id64)
		if errors.Is(err, storage.IdNotExistError) {
			logger.Printf("Path: / storageClient.Retrieve: %v\n", err)
			if err := return404(w); err != nil {
				logger.Printf("Path: / return404: %v\n", err)
			}
			return
		}
		if err != nil {
			logger.Printf("Path: / storageClient.Retrieve: %v\n", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, url, http.StatusFound)
	}
}
