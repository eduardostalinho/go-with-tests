package main

import "testing"

func TestEnglishGreet(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("says hello world for empty string", func(t *testing.T) {
		got := Greet("", "english")
		want := "Hello World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("says hello to harold", func(t *testing.T) {
		got := Greet("Harold", "english")
		want := "Hello Harold!"

		assertCorrectMessage(t, got, want)
	})
}

func TestSpanishGreet(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("says hola world for empty string", func(t *testing.T) {
		got := Greet("", "spanish")
		want := "Hola World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("says hola to harold", func(t *testing.T) {
		got := Greet("Harold", "spanish")
		want := "Hola Harold!"

		assertCorrectMessage(t, got, want)
	})
}

func TestGermanGreet(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("says hallo world for empty string", func(t *testing.T) {
		got := Greet("", "german")
		want := "Hallo World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("says hallo to harold", func(t *testing.T) {
		got := Greet("Harold", "german")
		want := "Hallo Harold!"

		assertCorrectMessage(t, got, want)
	})
}
