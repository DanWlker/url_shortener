package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"url_shortener/storage"
)

func ShortenHandler(logger *log.Logger) http.HandlerFunc {
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

		storage.MockStorage = append(storage.MockStorage, urlParam)
		// This should be in a async safe list that returns the id when appended, but whatever
		id := len(storage.MockStorage) - 1
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
		tmpl.Execute(w, struct{ Url string }{Url: fullShortenedUrl})
	}
}
