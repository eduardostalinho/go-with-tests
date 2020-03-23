package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello("")
	want := "Hello World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloHarold(t *testing.T) {
	got := Hello("Harold")
	want := "Hello Harold!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
