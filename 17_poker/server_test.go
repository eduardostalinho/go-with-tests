package poker_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"

	"github.com/eduardostalinho/go-with-tests/17_poker"
)

func TestGETPlayer(t *testing.T) {
	store := &poker.StubPlayerStore{
		Scores: map[string]int{
			"Adam":  20,
			"Alice": 10,
		},
	}
	server := mustMakePlayerServer(t, store, &poker.SpyGame{})
	t.Run("return Adam's score", func(t *testing.T) {
		request := newGetScoreRequest("Adam")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "20"

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), want)
	})

	t.Run("return Alice's score", func(t *testing.T) {
		request := newGetScoreRequest("Alice")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "10"

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), want)

	})

	t.Run("return not found for unexistent user", func(t *testing.T) {
		request := newGetScoreRequest("XPlayer")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		poker.AssertStatus(t, response, http.StatusNotFound)

	})
}

func TestPOSTPlayerScoresWins(t *testing.T) {
	store := &poker.StubPlayerStore{
		Scores: map[string]int{
			"Alice": 20,
			"Adam":  33,
		},
	}
	server := mustMakePlayerServer(t, store, &poker.SpyGame{})

	t.Run("returns and records wins", func(t *testing.T) {
		request := newPostScoreRequest("TestPlayer")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		poker.AssertStatus(t, response, http.StatusAccepted)
		poker.AssertResponseBody(t, response.Body.String(), "")
		poker.AssertPlayerWins(t, store, "TestPlayer")
	})
}

func TestLeague(t *testing.T) {
	league := poker.League{
		{"Alice", 10},
		{"Adam", 20},
	}
	store := &poker.StubPlayerStore{
		League: league,
	}
	server := mustMakePlayerServer(t, store, &poker.SpyGame{})

	t.Run("returns league sorted with highest scores first", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		wantedLeague := poker.League{
			{"Adam", 20},
			{"Alice", 10},
		}
		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
		contentType := response.Result().Header.Get("content-type")
		if contentType != "application/json" {
			t.Errorf("Expected content type to be application/json, got %s", contentType)
		}
		poker.AssertLeagueResponse(t, response.Body, wantedLeague)
	})
}

func TestGame(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		server := mustMakePlayerServer(t, store, &poker.SpyGame{})
		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
	})
	t.Run("starts with 3 players and the declare a winner", func(t *testing.T) {
		wantBlindAlert := "Blind is 100"
		winner := "Abel"

		game := &poker.SpyGame{BlindAlert: []byte(wantBlindAlert)}
		store := &poker.StubPlayerStore{}

		server := httptest.NewServer(mustMakePlayerServer(t, store, game))
		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws := mustDialWS(t, wsURL)

		defer server.Close()
		defer ws.Close()

		sendWSMessage(t, ws, "3")
		assertStartedCalledWith(t, game, 3)

		sendWSMessage(t, ws, winner)
		assertFinishCalledWith(t, game, winner)

		within(t, 10*time.Millisecond, func() {
			assertWebSocketGetMessage(t, ws, wantBlindAlert)
		})
	})
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func within(t *testing.T, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

func assertWebSocketGetMessage(t *testing.T, ws *websocket.Conn, want string) {
	_, msg, _ := ws.ReadMessage()
	if string(msg) != want {
		t.Errorf("expected message %s, got %s", want, string(msg))
	}
}

func assertStartedCalledWith(t *testing.T, game *poker.SpyGame, numberOfPlayers int) {
	t.Helper()
	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.NumberOfPlayers == numberOfPlayers
	})

	if !passed {
		t.Errorf("expected numberOfPlayers to be %d got %d", numberOfPlayers, game.NumberOfPlayers)
	}
}

func assertFinishCalledWith(t *testing.T, game *poker.SpyGame, winner string) {
	t.Helper()
	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.Winner == winner
	})

	if !passed {
		t.Errorf("expected game winner to be %s got %s", winner, game.Winner)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false

}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func newPlayerRequest(name, method string) *http.Request {
	request, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostScoreRequest(name string) *http.Request {
	return newPlayerRequest(name, http.MethodPost)
}

func newGetScoreRequest(name string) *http.Request {
	return newPlayerRequest(name, http.MethodGet)
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore, game *poker.SpyGame) *poker.PlayerServer {
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		t.Fatalf("could no start player server %v", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("Error trying to dial websocket on %s, %v", url, err)
	}
	return ws
}

func sendWSMessage(t *testing.T, conn *websocket.Conn, message string) {
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("Error trying to send message %v", err)
	}
}
