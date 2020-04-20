package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		got, err := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
		assertNoError(t, err)
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

		dictionary.Add("testa", "questa una testa")

		got, err := dictionary.Search("testa")
		want := "questa una testa"

		assertNoError(t, err)
		assertStrings(t, got, want)
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
