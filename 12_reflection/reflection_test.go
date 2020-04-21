package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Eduardo"
	var got []string
	x := struct {
		Name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	expectedCalls := 1
	if len(got) != expectedCalls {
		t.Errorf("wrong number of calls, got %d, want %d", len(got), expectedCalls)
	}
	if got[0] != expected {
		t.Errorf("got %q, want %q", got, expected)

	}

}
