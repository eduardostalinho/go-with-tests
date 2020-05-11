package poker

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "123456")
	defer clean()
	tape := &tape{file}

	newString := "abc"
	tape.Write([]byte(newString))
	file.Seek(0, 0)

	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)

	if got != newString {
		t.Errorf("got %q, want %q", got, newString)
	}
}

func TestFileSystemStore(t *testing.T) {
	t.Run("start store with empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()
		_, err := NewFileSystemStore(database)
		AssertNoError(t, err)
	})
	t.Run("get league", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Adam", "Wins": 20},
			{"Name": "Eve", "Wins": 10}
		]`)
		defer cleanDatabase()
		store, err := NewFileSystemStore(database)
		AssertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Adam", 20},
			{"Eve", 10},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

		got = store.GetLeague()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Adam", "Wins": 20},
			{"Name": "Eve", "Wins": 10}
		]`)
		defer cleanDatabase()
		store, err := NewFileSystemStore(database)
		AssertNoError(t, err)

		got := store.GetPlayerScore("Adam")

		want := 20
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("get player score for non-existent player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Adam", "Wins": 20},
			{"Name": "Eve", "Wins": 10}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		AssertNoError(t, err)

		got := store.GetPlayerScore("unexistent")

		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("get player score for non-existent player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Adam", "Wins": 20},
			{"Name": "Eve", "Wins": 10}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		AssertNoError(t, err)

		got := store.GetPlayerScore("NotExist")

		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("record win", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Adam", "Wins": 20},
			{"Name": "Eve", "Wins": 10}
		]`)
		defer cleanDatabase()
		store, err := NewFileSystemStore(database)
		AssertNoError(t, err)

		store.RecordWin("Adam")
		got := store.GetPlayerScore("Adam")

		want := 21
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("record win for new player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Adam", "Wins": 20},
			{"Name": "Eve", "Wins": 10}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		AssertNoError(t, err)

		store.RecordWin("Caim")
		got := store.GetPlayerScore("Caim")

		want := 1
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file. %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
