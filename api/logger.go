package main

import (
	"log"
	"net/http"
	"time"
)

// Logger is a decorator for a HTTP Request, adding logging functionality
func Logger(inner http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
