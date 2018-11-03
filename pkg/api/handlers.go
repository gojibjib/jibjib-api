package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// NotFound is a 404 Message according to the Response type
func NotFound(w http.ResponseWriter, r *http.Request) {
	SendErrorJSON(w, http.StatusInternalServerError, "Not found")
	return
}

// Ping sends a simple "Pong" as as JSON response
func (s *Server) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NewResponse(http.StatusOK, "Pong", 0, nil).SendJSON(w)
		return
	}
}

// Dummy sends a test Response mit randomized IDs and accuracies
func (s *Server) Dummy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id []int
		var acc []float64
		var data []Data

		seed := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(seed)
		n := 4

		// generate IDs and accs
		for i := 0; i < n; i++ {
			id = append(id, r1.Intn(194))
			acc = append(acc, r1.Float64()*100)
		}
		sort.Sort(sort.Reverse(sort.Float64Slice(acc)))

		// Generate data
		for i := 0; i < n; i++ {
			data = append(data, Data{id[i], acc[i]})
		}

		NewResponse(http.StatusOK, "Recognition successful", 4, data).SendJSON(w)
		return
	}
}

// GetAllBirds returns all Bird info from the database, without descriptions.
// This needs to be limited with more birds (200+).
func (s *Server) GetAllBirds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Session.Copy()
		defer session.Close()
		c := session.DB("birds").C("birds")

		var all []*Bird
		err := c.Find(bson.M{}).Select(bson.M{"desc_de": false, "desc_en": false}).All(&all)
		if err != nil {
			log.Println(err)
			SendErrorJSON(w, http.StatusInternalServerError, "Error retrieving list of birds")
			return
		}

		NewResponse(http.StatusOK, "Bird info retrieved", len(all), all).SendJSON(w)
		return
	}
}

// GetBirdByID Retrieves a single Bird from the database.
func (s *Server) GetBirdByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Session.Copy()
		defer session.Close()
		c := session.DB("birds").C("birds")

		// Parse arguments
		p := mux.Vars(r)
		id, err := strconv.Atoi(p["id"])
		if err != nil {
			SendErrorJSON(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		// Construct query
		descEN, descDE := r.FormValue("desc_en"), r.FormValue("desc_de")
		selections := make(bson.M)
		selec := false
		if descEN == "false" {
			selections["desc_en"] = false
			selec = true
		}
		if descDE == "false" {
			selections["desc_de"] = false
			selec = true
		}

		// Query DB & send response
		var bird *Bird
		q := c.Find(bson.M{"id": id})
		if selec {
			q = q.Select(selections)
		}
		err = q.One(&bird)
		if err != nil {
			log.Println(err)
			SendErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("Unable to find bird with id=%d", id))
			return
		}
		NewResponse(http.StatusOK, "Bird found", 1, bird).SendJSON(w)
		return
	}
}

// DetectBinary takes a application/octet-stream file via POST requests and forwards it to the API for detection
func (s *Server) DetectBinary() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read uploaded data
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			SendErrorJSON(w, http.StatusInternalServerError, "Error while trying to upload file")
			return
		}
		defer r.Body.Close()

		if len(b) <= 0 {
			SendErrorJSON(w, http.StatusBadRequest, "Uploaded file is empty")
			return
		}

		// Send data to Query Service
		reader := bytes.NewBuffer(b)
		//queryServiceURL := s.ModelURL + "/audio/transform/binary"
		queryServiceURL := s.ModelURL + "/detect/binary"
		contentType := "application/octet-stream"

		log.Printf("POST buffer (size: %d) to %s as %s\n", len(b), queryServiceURL, contentType)
		resp, err := http.Post(queryServiceURL, contentType, reader)
		if err != nil {
			log.Println(err)
			SendErrorJSON(w, http.StatusInternalServerError, "Error while trying to transmit file to query service")
			return
		}
		defer resp.Body.Close()

		// Parse query service response
		var a Response
		b, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			SendErrorJSON(w, http.StatusInternalServerError, "Error while trying to read query response")
			return
		}
		if err := json.Unmarshal(b, &a); err != nil {
			log.Println(err)
			SendErrorJSON(w, http.StatusInternalServerError, "Error while trying to unmarshal query response")
			return
		}

		NewResponse(a.Status, a.Message, a.Count, a.Data).SendJSON(w)
		return
	}
}
