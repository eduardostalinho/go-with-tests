package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayer(t *testing.T) {
	t.Run("return Adam's score", func(t *testing.T) {
		request := newGetScoreRequest("Adam")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		want := "20"

		assertResponseBody(t, response.Body.String(), want)
	})

	t.Run("return Alice's score", func(t *testing.T) {
		request := newGetScoreRequest("Alice")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		want := "10"

		assertResponseBody(t, response.Body.String(), want)

	})
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
