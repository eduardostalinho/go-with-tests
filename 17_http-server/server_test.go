package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayer(t *testing.T) {
	t.Run("return Adam's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Adam", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		want := "20"

		assertResponse(t, response, want)
	})

	t.Run("return Alice's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Alice", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		want := "10"

		assertResponse(t, response, want)

	})
}

func assertResponse(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := response.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
