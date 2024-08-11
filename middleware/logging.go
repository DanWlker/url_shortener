package middleware

import (
	"log"
	"net/http"
	"time"
)

type logHttpResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *logHttpResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			lhrw := &logHttpResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			next.ServeHTTP(lhrw, r)

			log.Println(lhrw.statusCode, r.Method, r.URL.Path, time.Since(start))
		},
	)
}
