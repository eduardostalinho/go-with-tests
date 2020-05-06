package main

import (
	"log"
	"net/http"
)

func main() {
	boltPath := "scores.bolt"
	bucketName := "scores"
	server := &PlayerServer{NewBoltPlayerStore(boltPath, bucketName)}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not start server on port 5000. Error: %v", err)
	}
}
