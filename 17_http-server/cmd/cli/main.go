package main

import (
	"fmt"
	"github.com/eduardostalinho/go-with-tests/17_http-server"
	"log"
	"os"
)

const dbFilename = "game.db.json"

func main() {
	fmt.Println("Let's play poker!")
	fmt.Println("type `{Name} wins` to record a win")

	db, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("unable to open database file %s, %v", dbFilename, err)
	}

	store, err := poker.NewFileSystemStore(db)

	if err != nil {
		log.Fatalf("unable to create file system store %v", err)
	}
	cli := poker.NewCLI(store, os.Stdin)
	cli.PlayPoker()
}
