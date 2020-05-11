package poker

import (
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRecordWinsAndRetrieveScore(t *testing.T) {
	cases := []struct {
		name          string
		serverFactory func(*testing.T) (*PlayerServer, func())
	}{
		{"with bolt", func(t *testing.T) (*PlayerServer, func()) {
			dbPath := "testdb.bolt"
			bucket := "scoresTest"
			store := NewBoltPlayerStore(dbPath, bucket)
			server := NewPlayerServer(store)
			return server, func() { tearDownBolt(store.bucketName, store.db) }
		}},
		{"with fs", func(t *testing.T) (*PlayerServer, func()) {
			database, cleanFsDatabase := createTempFile(t, "[]")
			store, err := NewFileSystemStore(database)
			if err != nil {
				t.Fatalf("unable to create file system store, %v", err)
			}
			server := NewPlayerServer(store)
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

			assertResponseBody(t, response.Body.String(), strconv.Itoa(wins))
		})
	}
}
