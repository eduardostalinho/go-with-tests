package osexec

import (
	"strings"
	"testing"
)

func TestIntegrationGetData(t *testing.T) {
	got := GetData(getXMLFromCommand())
	want := "HAPPY NEW YEAR!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetData(t *testing.T) {
	input := strings.NewReader(`
<payload>
	<message>Cats are the best!</message>
</payload>`)
	got := GetData(input)
	want := "CATS ARE THE BEST!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
