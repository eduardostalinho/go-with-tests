package contextawarereader

import (
	"strings"
	"testing"
)

func TestContextAwareReader(t *testing.T) {
	t.Run("normal reader", func(t *testing.T) {
		r := strings.NewReader("ABCDEF")
		got := make([]byte, 3)

		_, err := r.Read(got)

		abortOnError(t, err)

		assertBufferHas(t, got, "ABC")

		_, err = r.Read(got)
		abortOnError(t, err)

		assertBufferHas(t, got, "DEF")
	})
}

func abortOnError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertBufferHas(t *testing.T, buf []byte, want string) {
	t.Helper()
	got := string(buf)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
