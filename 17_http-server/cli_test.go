package poker_test

import (
	"bytes"
	"fmt"
	"github.com/eduardostalinho/go-with-tests/17_http-server"
	"strconv"
	"strings"
	"testing"
)

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
				game := &poker.SpyGame{}

				cli := poker.NewCLI(game, in, stdout)
				cli.PlayPoker()

				gotPrompt := stdout.String()
				if gotPrompt != poker.PlayerPrompt {
					t.Errorf("expected output %s, got %s", poker.PlayerPrompt, gotPrompt)
				}
				poker.AssertGame(t, game, numberOfPlayers, player)

			})
		}
	})
	t.Run("prints error on invalid number of players", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := userSends("Pies")
		game := &poker.SpyGame{}

		cli := poker.NewCLI(game, in, out)
		cli.PlayPoker()

		poker.AssertMessageSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputMsg)

		if game.StartCalled {
			t.Error("game should not have started")
		}
	})

	t.Run("prints error for invalid win message", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := userSends("3", "Lord is a killer")
		game := &poker.SpyGame{}

		cli := poker.NewCLI(game, in, out)
		cli.PlayPoker()

		poker.AssertMessageSentToUser(t, out, poker.PlayerPrompt, poker.BadWinMsg)

		if !game.StartCalled {
			t.Errorf("game should have started.")
		}

		if game.FinishCalled {
			t.Error("game should not have ended")
		}
	})
}
