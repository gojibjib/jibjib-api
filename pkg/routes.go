package jibjib_api

import "net/http"

//Route is a helper type defining a specific Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines a slice of all available API Routes
type Routes []Route

// Routes defines all HTTP routes, hanging off the main Server struct.
// Like that, all routes have access to the Server's dependencies.
func (s *Server) Routes() {
	var routes = Routes{
		Route{
			Name:        "Root",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: s.Ping(),
		},
		Route{
			Name:        "Ping",
			Method:      "GET",
			Pattern:     "/ping",
			HandlerFunc: s.Ping(),
		},
		Route{
			Name:        "DummyResponse",
			Method:      "GET",
			Pattern:     "/birds/dummy",
			HandlerFunc: s.Dummy(),
		},
		Route{
			Name:        "GetAllBirds",
			Method:      "GET",
			Pattern:     "/birds/all",
			HandlerFunc: s.GetAllBirds(),
		},
		Route{
			Name:        "GetBirdByID",
			Method:      "GET",
			Pattern:     "/birds/{id:[0-9]+}",
			HandlerFunc: s.GetBirdByID(),
		},
	}

	for _, route := range routes {
		var handler http.HandlerFunc
		handler = route.HandlerFunc
		handler = Logger(handler)
		s.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}
