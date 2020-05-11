package poker_test

import (
	"fmt"
	"github.com/eduardostalinho/go-with-tests/17_http-server"
	"strings"
	"testing"
	"time"
)

type alert struct {
	scheduledAt time.Duration
	amount      int
}
type SpyBlindAlerter struct {
	alerts []alert
}

func (a *SpyBlindAlerter) ScheduleAlert(duration time.Duration, amount int) {
	a.alerts = append(a.alerts, alert{duration, amount})
}

func TestCLI(t *testing.T) {
	players := []string{"Adam", "Eve"}
	for _, p := range players {
		t.Run(fmt.Sprintf("%s wins", p), func(t *testing.T) {
			in := strings.NewReader(fmt.Sprintf("%s wins\n", p))
			store := &poker.StubPlayerStore{}
			cli := poker.NewCLI(store, in, &SpyBlindAlerter{})
			cli.PlayPoker()

			poker.AssertPlayerWins(t, store, p)
		})
	}
	t.Run("schedules blind raise alerts", func(t *testing.T) {
		in := strings.NewReader("R2D2 wins\n")
		store := &poker.StubPlayerStore{}
		alerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(store, in, alerter)
		cli.PlayPoker()

		if len(alerter.alerts) != 1 {
			t.Fatal("expected blind alert to be scheduled")
		}
	})
}
