package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(1 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	fastURL := fastServer.URL
	slowURL := slowServer.URL

	defer slowServer.Close()
	defer fastServer.Close()

	got := Racer(fastURL, slowURL)
	want := fastURL

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
