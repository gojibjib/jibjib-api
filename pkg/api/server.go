package api

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Server defines holds information about routing and the Database sessions.
// Since all routes hang off the Server, handlers can easily access the database session.
type Server struct {
	Router  *mux.Router
	Session *mgo.Session
}
