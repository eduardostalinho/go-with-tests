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

type alert struct {
	scheduledAt time.Duration
	amount      int
}

func (a alert) String() string {
	return fmt.Sprintf("%d chips at %s", a.amount, a.scheduledAt)
}

type SpyBlindAlerter struct {
	alerts []alert
}

func (a *SpyBlindAlerter) ScheduleAlert(duration time.Duration, amount int) {
	a.alerts = append(a.alerts, alert{duration, amount})
}

func TestCLI(t *testing.T) {
	dummyAlerter := &SpyBlindAlerter{}
	dummyStore := &poker.StubPlayerStore{}
	dummyStdout := &bytes.Buffer{}

	t.Run("records players wins", func(t *testing.T) {
		players := []string{"Adam", "Eve"}
		for _, p := range players {
			t.Run(fmt.Sprintf("%s wins", p), func(t *testing.T) {
				in := strings.NewReader(fmt.Sprintf("5\n%s wins\n", p))
				store := &poker.StubPlayerStore{}
				cli := poker.NewCLI(store, in, dummyStdout, dummyAlerter)
				cli.PlayPoker()

				poker.AssertPlayerWins(t, store, p)
			})
		}
	})
	t.Run("schedules printing of values", func(t *testing.T) {
		in := strings.NewReader("5\n")
		alerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(dummyStore, in, dummyStdout, alerter)
		cli.PlayPoker()

		cases := []alert{
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
			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
			}
			scheduledAlert := alerter.alerts[i]
			assertScheduledAlert(t, c, scheduledAlert)
		}
	})
	t.Run("prompts for input of number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		alerter := &SpyBlindAlerter{}
		in := strings.NewReader("7\n")

		cli := poker.NewCLI(dummyStore, in, stdout, alerter)
		cli.PlayPoker()

		got := stdout.String()
		want := "Please input the number of players: "

		if got != want {
			t.Errorf("got stdout %q, want %q", got, want)
		}

		cases := []alert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}
		for i, c := range cases {
			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
			}
			scheduledAlert := alerter.alerts[i]
			assertScheduledAlert(t, c, scheduledAlert)
		}

	})
}

func assertScheduledAlert(t *testing.T, alertWanted, scheduledAlert alert) {
	if !reflect.DeepEqual(alertWanted, scheduledAlert) {
		t.Errorf("expected alert %v got %v", alertWanted, scheduledAlert)
	}
}
