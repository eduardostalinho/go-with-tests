package main

import (
	"github.com/eduardostalinho/go-with-tests/17_http-server"
	"log"
	"net/http"
)

func main() {
	boltPath := "scores.bolt"
	bucketName := "scores"
	store := poker.NewBoltPlayerStore(boltPath, bucketName)

	game := poker.NewGame(store, poker.BlindAlerterFunc(poker.Alerter))
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("Cound not start game server.")
	}

	defer store.Close()
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not start server on port 5000. Error: %v", err)
	}
}
