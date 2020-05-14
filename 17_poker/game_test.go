package poker_test

import (
	"io/ioutil"
	"reflect"
	"testing"
	"time"

	"github.com/eduardostalinho/go-with-tests/17_poker"
)

func TestTexasHoldEM_Start(t *testing.T) {
	dummyStore := &poker.StubPlayerStore{}
	t.Run("starts game for 7 players", func(t *testing.T) {
		alerter := &poker.SpyBlindAlerter{}
		numberOfPlayers := 7

		game := poker.NewTexasHoldEm(dummyStore, alerter)
		game.Start(numberOfPlayers, ioutil.Discard)

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
	t.Run("schedules printing of values for 5 players", func(t *testing.T) {
		numberOfPlayers := 5
		alerter := &poker.SpyBlindAlerter{}

		game := poker.NewTexasHoldEm(dummyStore, alerter)
		game.Start(numberOfPlayers, ioutil.Discard)

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
}

func assertScheduledAlert(t *testing.T, alertWanted, scheduledAlert poker.SpyAlert) {
	if !reflect.DeepEqual(alertWanted, scheduledAlert) {
		t.Errorf("expected alert %v got %v", alertWanted, scheduledAlert)
	}
}
