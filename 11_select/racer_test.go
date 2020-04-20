package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("selects faster server", func(t *testing.T) {
		slowServer := makeDelayedServer(1 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		got, err := Racer(fastURL, slowURL)
		want := fastURL

		assertNoError(t, err)
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("raiser error if server dont respond in 10 seconds", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Second)
		fastServer := makeDelayedServer(12 * time.Second)

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer(fastURL, slowURL)

		if err != ErrTimeoutErr {
			t.Errorf("error mismatch, got %q", err)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("unexpected error %q", err)
	}
}
