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
		"Root",
		"GET",
		"/",
		Ping,
	},
	Route{
		"Ping",
		"GET",
		"/ping",
		Ping,
	},
	Route{
		Name:        "Dummy response",
		Method:      "GET",
		Pattern:     "/dummy",
		HandlerFunc: Dummy,
	},
}
