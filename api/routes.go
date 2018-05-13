package main

import "net/http"

// Route is a helper type defining a specific Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines a slice of all available API Routes
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Root",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Ping,
	},
	Route{
		Name:        "Ping",
		Method:      "GET",
		Pattern:     "/ping",
		HandlerFunc: Ping,
	},
	Route{
		Name:        "Dummy response",
		Method:      "GET",
		Pattern:     "/dummy",
		HandlerFunc: Dummy,
	},
}
