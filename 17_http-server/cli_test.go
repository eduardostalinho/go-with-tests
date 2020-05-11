package poker_test

import (
	"bytes"
	"fmt"
	"github.com/eduardostalinho/go-with-tests/17_http-server"
	"strconv"
	"strings"
	"testing"
)

type SpyGame struct {
	numberOfPlayers int
	winner          string
	startCalled     bool
	finishCalled    bool
}

func (g *SpyGame) Start(numberOfPlayers int) {
	g.startCalled = true
	g.numberOfPlayers = numberOfPlayers
}

func (g *SpyGame) Finish(winner string) {
	g.finishCalled = true
	g.winner = winner
}

func TestCLI(t *testing.T) {
	userSends := func(inputs ...string) *strings.Reader {
		messages := strings.Join(inputs, "\n") + "\n"
		return strings.NewReader(messages)

	}
	t.Run("finish game from user input", func(t *testing.T) {
		players := []string{"Adam", "Eve"}
		for _, player := range players {
			t.Run(fmt.Sprintf("%s wins", player), func(t *testing.T) {
				numberOfPlayers := 5

				stdout := &bytes.Buffer{}
				in := userSends(strconv.Itoa(numberOfPlayers), fmt.Sprintf("%s wins", player))
				game := &SpyGame{}

				cli := poker.NewCLI(game, in, stdout)
				cli.PlayPoker()

				gotPrompt := stdout.String()
				if gotPrompt != poker.PlayerPrompt {
					t.Errorf("expected output %s, got %s", poker.PlayerPrompt, gotPrompt)
				}
				assertGame(t, game, numberOfPlayers, player)

			})
		}
	})
	t.Run("prints error on invalid number of players", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := userSends("Pies")
		game := &SpyGame{}

		cli := poker.NewCLI(game, in, out)
		cli.PlayPoker()

		assertMessageSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputMsg)

		if game.startCalled {
			t.Error("game should not have started")
		}
	})

	t.Run("prints error for invalid win message", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := userSends("3", "Lord is a killer")
		game := &SpyGame{}

		cli := poker.NewCLI(game, in, out)
		cli.PlayPoker()

		assertMessageSentToUser(t, out, poker.PlayerPrompt, poker.BadWinMsg)

		if !game.startCalled {
			t.Errorf("game should have started.")
		}

		if game.finishCalled {
			t.Error("game should not have ended")
		}
	})
}

func assertMessageSentToUser(t *testing.T, out *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := out.String()

	if got != want {
		t.Fatalf("got out %q, want %q", got, want)
	}
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
