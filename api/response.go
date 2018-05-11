package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

func NewResponse(m string, d []string) Response {
	return Response{
		Status:  http.StatusOK,
		Message: m,
		Data:    d,
	}
}

func (r Response) sendJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/JSON; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		status := http.StatusInternalServerError
		w.WriteHeader(status)

		r.Status = status
		r.Message = "Internal Server Error"
		r.Data = nil

		_ = json.NewEncoder(w).Encode(r)
	}
}
