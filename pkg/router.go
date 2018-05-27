package jibjib_api

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter returns a gorilla/mux router with a custom 404 handler.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFound)
	return router
}
