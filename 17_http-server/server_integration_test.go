package poker_test

import (
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/eduardostalinho/go-with-tests/17_http-server"
)

func TestRecordWinsAndRetrieveScore(t *testing.T) {
	cases := []struct {
		name          string
		serverFactory func(*testing.T) (*poker.PlayerServer, func())
	}{
		{"with bolt", func(t *testing.T) (*poker.PlayerServer, func()) {
			dbPath := "testdb.bolt"
			bucket := "scoresTest"
			store := poker.NewBoltPlayerStore(dbPath, bucket)
			server, err := poker.NewPlayerServer(store, &poker.SpyGame{})
			if err != nil {
				t.Fatalf("unexpected error instantiating server %v", err)
			}
			return server, func() { tearDownBolt(store.BucketName, store.DB) }
		}},
		{"with fs", func(t *testing.T) (*poker.PlayerServer, func()) {
			database, cleanFsDatabase := poker.CreateTempFile(t, "[]")
			store, err := poker.NewFileSystemStore(database)
			if err != nil {
				t.Fatalf("unable to create file system store, %v", err)
			}
			server, err := poker.NewPlayerServer(store, &poker.SpyGame{})
			if err != nil {
				t.Fatalf("unexpected error instantiating server %v", err)
			}
			return server, func() { cleanFsDatabase() }
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			server, closeDB := c.serverFactory(t)
			defer closeDB()

			player := "Agnawd"
			wins := 3

			for i := 0; i < wins; i++ {
				server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
			}

			response := httptest.NewRecorder()
			server.ServeHTTP(response, newGetScoreRequest(player))

			poker.AssertResponseBody(t, response.Body.String(), strconv.Itoa(wins))
		})
	}
}
