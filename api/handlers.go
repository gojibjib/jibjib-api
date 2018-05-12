package main

import (
	"log"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse("Pong", nil)
	resp.sendJSON(w)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	resp := NewResponse("Not found", nil)
	resp.Status = http.StatusNotFound
	log.Printf(
		"%s\t%s\t%s\t%s",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
	resp.sendJSON(w)
}

func Dummy(w http.ResponseWriter, r *http.Request) {
	var id []int
	var acc []float64
	var data []Data

	seed := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(seed)
	n := 4

	// generate IDs and accs
	for i := 0; i < n; i++ {
		id = append(id, r1.Intn(1190))
		acc = append(acc, r1.Float64()*100)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(acc)))

	// Generate data
	for i := 0; i < n; i++ {
		data = append(data, Data{id[i], acc[i]})
	}

	resp := NewResponse("Recognition successful", data)
	resp.sendJSON(w)
}
