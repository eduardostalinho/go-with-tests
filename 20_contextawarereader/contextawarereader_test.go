package contextawarereader

import (
	"context"
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
	t.Run("context-aware reader behaves like normal reader", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		r := NewCancellableReader(ctx, strings.NewReader("ABCDEF"))
		got := make([]byte, 3)

		_, err := r.Read(got)

		abortOnError(t, err)

		assertBufferHas(t, got, "ABC")

		_, err = r.Read(got)
		abortOnError(t, err)

		assertBufferHas(t, got, "DEF")
		cancel()
	})
	t.Run("context-aware reader stops reading on cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		r := NewCancellableReader(ctx, strings.NewReader("ABCDEF"))
		got := make([]byte, 3)

		_, err := r.Read(got)

		abortOnError(t, err)

		assertBufferHas(t, got, "ABC")

		cancel()

		n, err := r.Read(got)

		expectError(t, err)

		if n > 0 {
			t.Errorf("expected nothing read, but read %d bytes", n)
		}
	})
}

func expectError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected an error")
	}

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
