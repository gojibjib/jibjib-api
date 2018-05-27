package main

import (
	"github.com/gojibjib/jibjib-api/pkg/api"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
)

func main() {
	env := "JIBJIB_DB_URL"
	mgoURL := os.Getenv(env)
	if mgoURL == "" {
		log.Fatalf("Environment variable %s is empty", env)
	}

	session, err := mgo.Dial(mgoURL)
	if err != nil {
		log.Fatalf("Unable to establish DB connection to %s: %s", mgoURL, err)
	}
	defer session.Close()
	log.Printf("DB connection established at %s", mgoURL)

	server := api.Server{
		Router:  api.NewRouter(),
		Session: session,
	}

	server.Routes()
	addr := "0.0.0.0:8080"
	log.Printf("Starting listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
