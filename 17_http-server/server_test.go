package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
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
func TestGETPlayer(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Adam":  20,
			"Alice": 10,
		},
	}
	server := NewPlayerServer(&store)
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
			"Alice": 20,
			"Adam":  33,
		},
	}
	server := NewPlayerServer(&store)

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

func TestLeague(t *testing.T) {
	league := League{
		{"Alice", 10},
		{"Adam", 20},
	}
	store := StubPlayerStore{
		league: league,
	}
	server := NewPlayerServer(&store)

	t.Run("returns league sorted with highest scores first", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		wantedLeague := League{
			{"Adam", 20},
			{"Alice", 10},
		}
		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		contentType := response.Result().Header.Get("content-type")
		if contentType != "application/json" {
			t.Errorf("Expected content type to be application/json, got %s", contentType)
		}
		assertLeagueResponse(t, response.Body, wantedLeague)
	})
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

func assertLeagueResponse(t *testing.T, body io.Reader, want []Player) {
	got, _ := LeagueFromReader(body)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
