package poker

import (
	"fmt"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

type SpyAlert struct {
	ScheduledAt time.Duration
	Amount      int
}

func (a SpyAlert) String() string {
	return fmt.Sprintf("%d chips at %s", a.Amount, a.ScheduledAt)
}

type SpyBlindAlerter struct {
	Alerts []SpyAlert
}

func (a *SpyBlindAlerter) ScheduleAlert(duration time.Duration, amount int) {
	a.Alerts = append(a.Alerts, SpyAlert{duration, amount})
}

func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertStatus(t *testing.T, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	got := response.Code
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertLeagueResponse(t *testing.T, body io.Reader, want []Player) {
	t.Helper()
	got, _ := LeagueFromReader(body)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func AssertPlayerWins(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) != 1 {
		t.Fatal("expected win called once")
	}
	got := store.winCalls[0]

	if got != winner {
		t.Errorf("expected win %s, got %s", winner, got)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
}
