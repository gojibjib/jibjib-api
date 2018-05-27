package main

import (
	"github.com/gojibjib/jibjib-api/pkg"
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
	log.Printf("DB connection established at %s", mgoURL)

	server := jibjib_api.Server{
		Router:  jibjib_api.NewRouter(),
		Session: session,
	}

	server.Routes()
	log.Printf("Starting listening on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", server.Router))

}
