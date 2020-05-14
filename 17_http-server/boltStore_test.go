package poker_test

import (
	"reflect"
	"strconv"
	"testing"

	bolt "go.etcd.io/bbolt"

	"github.com/eduardostalinho/go-with-tests/17_http-server"
)

func setupBolt(path, bucket string) *bolt.DB {
	db, _ := bolt.Open(path, 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
	return db
}

func tearDownBolt(bucket string, db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucket))
		return err
	})
	db.Close()
}

func TestBoltStoreRecordWin(t *testing.T) {
	dbPath := "testdb.bolt"
	bucket := "scoresTest"
	db := setupBolt(dbPath, bucket)
	defer tearDownBolt(bucket, db)

	store := poker.BoltPlayerStore{db, bucket}

	t.Run("create new player with score 1", func(t *testing.T) {
		store.RecordWin("TestPlayer")
		got := make(chan int, 1)
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("scoresTest"))
			v := b.Get([]byte("TestPlayer"))
			res, _ := strconv.Atoi(string(v))
			got <- res
			return nil
		})

		res := <-got
		if res != 1 {
			t.Errorf("Expected value to be 1, got %d", res)
		}

	})

	t.Run("increment score by 1 for exisiting player", func(t *testing.T) {
		store.RecordWin("ExistingPlayer")
		store.RecordWin("ExistingPlayer")
		got := make(chan int, 1)
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("scoresTest"))
			v := b.Get([]byte("ExistingPlayer"))
			res, _ := strconv.Atoi(string(v))
			got <- res
			return nil
		})

		res := <-got
		if res != 2 {
			t.Errorf("Expected value to be 2, got %d", res)
		}

	})
}

func TestBoltStoreGetPlayerScore(t *testing.T) {
	dbPath := "testdb.bolt"
	bucket := "scoresTest"
	db := setupBolt(dbPath, bucket)
	defer tearDownBolt(bucket, db)

	store := poker.BoltPlayerStore{db, "scoresTest"}

	t.Run("get score for exisiting player", func(t *testing.T) {
		store.RecordWin("TestPlayer")
		got := store.GetPlayerScore("TestPlayer")

		if got != 1 {
			t.Errorf("Expected value to be 1, got %d", got)
		}
	})

	t.Run("get score 0 for non exisiting player", func(t *testing.T) {
		got := store.GetPlayerScore("NewPlayer")

		if got != 0 {
			t.Errorf("Expected value to be 1, got %d", got)
		}
	})
}

func TestBoltStoreGetLeague(t *testing.T) {
	dbPath := "testdb.bolt"
	bucket := "scoresTest"
	db := setupBolt(dbPath, bucket)
	defer tearDownBolt(bucket, db)

	store := poker.BoltPlayerStore{db, "scoresTest"}

	t.Run("get score for all players", func(t *testing.T) {
		store.RecordWin("TestPlayer")
		store.RecordWin("TestPlayer2")
		store.RecordWin("TestPlayer3")
		got := store.GetLeague()

		want := poker.League{
			{"TestPlayer", 1},
			{"TestPlayer2", 1},
			{"TestPlayer3", 1},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("get score 0 for non exisiting player", func(t *testing.T) {
		got := store.GetPlayerScore("NewPlayer")

		if got != 0 {
			t.Errorf("Expected value to be 0, got %d", got)
		}
	})
}
