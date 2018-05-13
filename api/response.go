package main

import (
	"encoding/json"
	"net/http"
)

// Response defines an API response.
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Data `json:"data"`
}

// Data defines the data field in the API response.
type Data struct {
	ID       int     `json:"id"`
	Accuracy float64 `json:"accuracy"`
}

// NewResponse returns a Response with a passed message string and slice of Data.
// This will automatically set the Status field to 200.
func NewResponse(m string, d []Data) Response {
	return Response{
		Status:  http.StatusOK,
		Message: m,
		Data:    d,
	}
}

// sendJSON encodes a Response as JSON and sends it on a passed http.ResponseWriter.
func (r Response) sendJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/JSON; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		status := http.StatusInternalServerError
		w.WriteHeader(status)

		r.Status = status
		r.Message = "Internal Server Error while trying to encode JSON"
		r.Data = nil

		_ = json.NewEncoder(w).Encode(r)
	}
}
