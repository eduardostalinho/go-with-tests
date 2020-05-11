package poker_test

import (
	"bytes"
	"fmt"
	"github.com/eduardostalinho/go-with-tests/17_http-server"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	dummyAlerter := &poker.SpyBlindAlerter{}
	dummyStore := &poker.StubPlayerStore{}
	dummyStdout := &bytes.Buffer{}

	t.Run("records players wins", func(t *testing.T) {
		players := []string{"Adam", "Eve"}
		for _, p := range players {
			t.Run(fmt.Sprintf("%s wins", p), func(t *testing.T) {
				in := strings.NewReader(fmt.Sprintf("5\n%s wins\n", p))
				store := &poker.StubPlayerStore{}
				game := poker.NewGame(store, dummyAlerter)
				cli := poker.NewCLI(game, in, dummyStdout)
				cli.PlayPoker()

				poker.AssertPlayerWins(t, store, p)
			})
		}
	})
	t.Run("schedules printing of values", func(t *testing.T) {
		in := strings.NewReader("5\n")
		alerter := &poker.SpyBlindAlerter{}

		game := poker.NewGame(dummyStore, alerter)
		cli := poker.NewCLI(game, in, dummyStdout)
		cli.PlayPoker()

		cases := []poker.SpyAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}
		for i, c := range cases {
			if len(alerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.Alerts)
			}
			scheduledAlert := alerter.Alerts[i]
			assertScheduledAlert(t, c, scheduledAlert)
		}
	})
	t.Run("prompts for input of number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		alerter := &poker.SpyBlindAlerter{}
		in := strings.NewReader("7\n")

		game := poker.NewGame(dummyStore, alerter)
		cli := poker.NewCLI(game, in, stdout)
		cli.PlayPoker()

		got := stdout.String()
		want := "Please input the number of players: "

		if got != want {
			t.Errorf("got stdout %q, want %q", got, want)
		}

		cases := []poker.SpyAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}
		for i, c := range cases {
			if len(alerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.Alerts)
			}
			scheduledAlert := alerter.Alerts[i]
			assertScheduledAlert(t, c, scheduledAlert)
		}

	})
}

func assertScheduledAlert(t *testing.T, alertWanted, scheduledAlert poker.SpyAlert) {
	if !reflect.DeepEqual(alertWanted, scheduledAlert) {
		t.Errorf("expected alert %v got %v", alertWanted, scheduledAlert)
	}
}
