package main

import (
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
		log.Fatalf("Unable to establish DB connection to %s", mgoURL)
	}
	defer session.Close()

	server := Server{
		Router:  NewRouter(),
		Session: session,
	}

	server.Routes()
	log.Fatal(http.ListenAndServe(":8080", server.Router))
}
