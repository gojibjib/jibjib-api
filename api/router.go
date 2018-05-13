package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter returns a gorilla/mux router with routes registered
// from routes.go
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFound)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
