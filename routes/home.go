package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
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

func NewHomeHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle 404
		if r.URL.Path != "/" {
			if err := return404(w); err != nil {
				logger.Printf("Path: / return404: %v\n", err)
			}
			return
		}

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
	}
}
