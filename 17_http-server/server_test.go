package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
}
func TestGETPlayer(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Adam":  20,
			"Alice": 10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("return Adam's score", func(t *testing.T) {
		request := newGetScoreRequest("Adam")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "20"

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), want)
	})

	t.Run("return Alice's score", func(t *testing.T) {
		request := newGetScoreRequest("Alice")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "10"

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), want)

	})

	t.Run("return not found for unexistent user", func(t *testing.T) {
		request := newGetScoreRequest("XPlayer")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseStatus(t, response.Code, http.StatusNotFound)

	})
}

func TestPOSTPlayerScoresWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Adam":  20,
			"Alice": 10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns accepted", func(t *testing.T) {
		request := newPostScoreRequest("TestPlayer")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseStatus(t, response.Code, http.StatusAccepted)
		assertResponseBody(t, response.Body.String(), "")
		if len(store.winCalls) != 1 {
			t.Errorf("expected RecordWin to be called 1 time., called %d", len(store.winCalls))
		}
	})
}

func newPostScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
