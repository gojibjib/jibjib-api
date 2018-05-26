package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// Ping sends a simple "Pong" as as JSON response
func Ping(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse("Pong", nil, http.StatusOK)
	resp.SendJSON(w)
}

// NotFound is a 404 Message according to the Response type
func NotFound(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse("Not found", nil, http.StatusNotFound)
	resp.SendJSON(w)
}

// Dummy sends a test Response mit randomized IDs and accuracies
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

	resp := NewResponse("Recognition successful", data, http.StatusOK)
	resp.SendJSON(w)
}

func GetBirdInfo(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	id := r.FormValue("id")
	if m, err := regexp.MatchString("[0-9]+", id); !m || err != nil {
		NewResponse("Invalid ID", nil, http.StatusBadRequest).SendJSON(w)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		NewResponse("Invalid ID", nil, http.StatusBadRequest).SendJSON(w)
		return
	}

	// Establish MongoDB connection
	mongoUrl := ""
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		log.Printf("Unable to establish DB connection to %s", mongoUrl)
		NewResponse("Unable to establish database connection", nil, http.StatusInternalServerError).SendJSON(w)
		return
	}
	defer session.Close()

	c := session.DB("birds").C("birds")

	var bird *Bird
	err = c.Find(bson.M{"id": idInt}).One(&bird)
	if err != nil {
		NewResponse(fmt.Sprintf("Unable to find bird by id=%d", idInt), nil, http.StatusNotFound).SendJSON(w)
		return
	}

	resp := NewResponse("Bird info found", bird, http.StatusOK)
	resp.SendJSON(w)
	return
}
