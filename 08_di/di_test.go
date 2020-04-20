package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Eduardo")

	got := buffer.String()
	want := "Hello Eduardo!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
