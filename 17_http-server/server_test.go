package poker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

		AssertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), want)
	})

	t.Run("return Alice's score", func(t *testing.T) {
		request := newGetScoreRequest("Alice")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "10"

		AssertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), want)

	})

	t.Run("return not found for unexistent user", func(t *testing.T) {
		request := newGetScoreRequest("XPlayer")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertStatus(t, response, http.StatusNotFound)

	})
}

func TestPOSTPlayerScoresWins(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Alice": 20,
			"Adam":  33,
		},
	}
	server := NewPlayerServer(store)

	t.Run("returns and records wins", func(t *testing.T) {
		request := newPostScoreRequest("TestPlayer")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertStatus(t, response, http.StatusAccepted)
		AssertResponseBody(t, response.Body.String(), "")
		AssertPlayerWins(t, store, "TestPlayer")
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

		AssertStatus(t, response, http.StatusOK)
		contentType := response.Result().Header.Get("content-type")
		if contentType != "application/json" {
			t.Errorf("Expected content type to be application/json, got %s", contentType)
		}
		AssertLeagueResponse(t, response.Body, wantedLeague)
	})
}

func TestGame(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		server := NewPlayerServer(&StubPlayerStore{})
		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response, http.StatusOK)
	})
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
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
