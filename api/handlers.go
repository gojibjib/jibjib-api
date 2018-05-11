package main

import (
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse("Pong", nil)
	resp.sendJSON(w)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse("Not found", nil)
	resp.Status = http.StatusNotFound
	resp.sendJSON(w)
}
