package main

import (
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRecordWinsAndRetrieveScore(t *testing.T) {
	dbPath := "testdb.bolt"
	bucket := "scoresTest"

	store := NewBoltPlayerStore(dbPath, bucket)
	defer store.db.Close()
	t.Run("record 3 wins and get score", func(t *testing.T) {
		server := &PlayerServer{store}
		player := "Agnawd"
		wins := 3

		for i := 0; i < wins; i++ {
			server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
		}

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertResponseBody(t, response.Body.String(), strconv.Itoa(wins))
	})
}
