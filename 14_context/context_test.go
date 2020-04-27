package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(10 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertCancelled(t *testing.T) {
	t.Helper()
	if !s.cancelled {
		t.Errorf("store should be cancelled")
	}
}

func (s *SpyStore) assertNotCancelled(t *testing.T) {
	t.Helper()
	if s.cancelled {
		t.Errorf("store should not be cancelled")
	}
}

func TestHandler(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, wanted %s", response.Body.String(), data)
		}

		store.assertNotCancelled(t)

	})

	t.Run("test store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"

		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertCancelled(t)
	})
}
