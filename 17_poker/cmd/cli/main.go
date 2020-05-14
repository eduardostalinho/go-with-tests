package main

import (
	"fmt"
	"github.com/eduardostalinho/go-with-tests/17_poker"
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

	alerter := poker.BlindAlerterFunc(poker.Alerter)
	game := poker.NewGame(store, alerter)
	cli := poker.NewCLI(game, os.Stdin, os.Stdout)
	cli.PlayPoker()
}
