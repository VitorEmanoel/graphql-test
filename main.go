package main

import (
	log "github.com/sirupsen/logrus"
	"graphql-test/api"
	"net/http"
)

func main() {
	api.NewGraphQL()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.WithField("error", err.Error()).Error("Error in start HTTP Server")
	}
}
