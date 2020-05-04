package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertExpectedResult := func(t *testing.T, got, expected string) {
		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	}
	t.Run("Repeat a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertExpectedResult(t, repeated, expected)
	})

	t.Run("Repeat b 5 times", func(t *testing.T) {
		repeated := Repeat("b", 5)
		expected := "bbbbb"

		assertExpectedResult(t, repeated, expected)
	})

	t.Run("Repeat a 6 times", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"

		assertExpectedResult(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}

}
