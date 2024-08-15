package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/DanWlker/url_shortener/storage"
)

func ShortenHandler(logger *log.Logger, storageClient storage.StorageClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			logger.Printf("Path: /shorten r.ParseForm: %v\n", err)
			return
		}

		urlParam := r.Form.Get("url")
		if urlParam == "" {
			http.Error(w, "Missing url", http.StatusBadRequest)
			logger.Printf("Path: /shorten Form.Get: Missing 'url' parameter \n")
			return
		}

		id, err := storageClient.Insert(urlParam)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			logger.Printf("storage.StorageClient.Insert: %v\n", err)
			return
		}

		idHex := fmt.Sprintf("%x", id)

		fullShortenedUrl, err := url.JoinPath("http://", r.Host, idHex)
		fmt.Println(fullShortenedUrl)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			logger.Printf("Path: /shorten url.JoinPath: %v\n", err)
			return
		}

		tmpl, err := template.ParseFiles("templates/shorten.html")
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			logger.Printf("Path: /shorten template.ParseFiles: %v\n", err)
			return
		}
		if err := tmpl.Execute(w, struct{ Url string }{Url: fullShortenedUrl}); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			logger.Printf("Path: /shorten tmpl.Execute: %v\n", err)
			return
		}
	}
}
