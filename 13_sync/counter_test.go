package counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("3 Inc calls make value 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}

func assertCounter(t *testing.T, counter Counter, value int) {
	t.Helper()
	if counter.Value() != value {
		t.Errorf("got %d, wanted %d", counter.Value(), value)
	}
}
