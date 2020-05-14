package poker

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

type SpyGame struct {
	numberOfPlayers int
	winner          string
	StartCalled     bool
	FinishCalled    bool
	BlindAlert      []byte
}

func (g *SpyGame) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.numberOfPlayers = numberOfPlayers
	out.Write(g.BlindAlert)
}

func (g *SpyGame) Finish(winner string) {
	g.FinishCalled = true
	g.winner = winner
}

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.Scores[player]
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.WinCalls = append(s.WinCalls, player)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
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

func (a *SpyBlindAlerter) ScheduleAlert(duration time.Duration, amount int, to io.Writer) {
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
	if len(store.WinCalls) != 1 {
		t.Fatal("expected win called once")
	}
	got := store.WinCalls[0]

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

func AssertMessageSentToUser(t *testing.T, out *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := out.String()

	if got != want {
		t.Fatalf("got out %q, want %q", got, want)
	}
}

func AssertGame(t *testing.T, game *SpyGame, numberOfPlayers int, winner string) {
	t.Helper()
	if game.numberOfPlayers != numberOfPlayers {
		t.Errorf("expected number of players %d, got %d", numberOfPlayers, game.numberOfPlayers)
	}

	if game.winner != winner {
		t.Errorf("expected winner %s, got %s", winner, game.winner)
	}
}

func CreateTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file. %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
