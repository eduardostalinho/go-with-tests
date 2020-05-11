package poker_test

import (
	"bytes"
	"fmt"
	"github.com/eduardostalinho/go-with-tests/17_http-server"
	"strings"
	"testing"
)

type SpyGame struct {
	numberOfPlayers int
	winner          string
}

func (g *SpyGame) Start(numberOfPlayers int) {
	g.numberOfPlayers = numberOfPlayers
}

func (g *SpyGame) Finish(winner string) {
	g.winner = winner
}

func TestCLI(t *testing.T) {

	t.Run("finish game from user input", func(t *testing.T) {
		players := []string{"Adam", "Eve"}
		for _, p := range players {
			t.Run(fmt.Sprintf("%s wins", p), func(t *testing.T) {
				numberOfPlayers := 5

				stdout := &bytes.Buffer{}
				in := strings.NewReader(fmt.Sprintf("%d\n%s wins\n", numberOfPlayers, p))
				game := &SpyGame{}

				cli := poker.NewCLI(game, in, stdout)
				cli.PlayPoker()

				gotPrompt := stdout.String()
				if gotPrompt != poker.PlayerPrompt {
					t.Errorf("expected output %s, got %s", poker.PlayerPrompt, gotPrompt)
				}
				assertGame(t, game, numberOfPlayers, p)

			})
		}
	})
}

func assertGame(t *testing.T, game *SpyGame, numberOfPlayers int, winner string) {
	t.Helper()
	if game.numberOfPlayers != numberOfPlayers {
		t.Errorf("expected number of players %d, got %d", numberOfPlayers, game.numberOfPlayers)
	}

	if game.winner != winner {
		t.Errorf("expected winner %s, got %s", winner, game.winner)
	}
}
