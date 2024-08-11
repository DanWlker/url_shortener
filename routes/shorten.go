package routes

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func ShortenHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			logger.Printf("Path: /shorten io.ReadAll: %v\n", err)
			return
		}

		logger.Printf("POST params were: %s\n", b)

		tmpl, err := template.ParseFiles("templates/shorten.html")
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			logger.Printf("Path: /shorten template.ParseFiles: %v\n", err)
			return
		}
		tmpl.Execute(w, struct{ Url string }{Url: "hello"})
	}
}
