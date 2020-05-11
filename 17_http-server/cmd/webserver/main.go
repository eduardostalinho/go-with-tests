package main

import (
	"log"
	"net/http"
	"github.com/eduardostalinho/go-with-tests/17_http-server"
)

func main() {
	boltPath := "scores.bolt"
	bucketName := "scores"
	store := poker.NewBoltPlayerStore(boltPath, bucketName)
	server := poker.NewPlayerServer(store)
	defer store.Close()
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not start server on port 5000. Error: %v", err)
	}
}
