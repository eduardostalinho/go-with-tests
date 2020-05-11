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

	store, close, err := poker.FileSystemStoreFromFile(dbFilename)
	if err != nil {
		log.Fatal(err)

	}
	defer close()

	cli := poker.NewCLI(store, os.Stdin, poker.BlindAlerterFunc(poker.StdoutAlerter))
	cli.PlayPoker()
}
