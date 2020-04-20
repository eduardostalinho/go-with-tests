package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}

		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		_, err := dictionary.Search("tests")
		assertError(t, err, ErrUnexistingWord)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add unexisting word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		word := "testa"
		definition := "questa una testa"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t *testing.T, err, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected error")
	}

	if err != want {
		t.Errorf("error mismatch: got %q want %q", err, want)
	}

}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("unexpected error %q", err)
	}
}

func assertDefinition(t *testing.T, d Dictionary, w, definition string) {
	t.Helper()
	got, err := d.Search(w)

	assertNoError(t, err)
	assertStrings(t, got, definition)
}
