package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("3 Inc calls make value 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	t.Run("Increment asynchronously 1000 times", func(t *testing.T) {
		counter := NewCounter()
		wantedCount := 1000
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, 1000)
	})
}

func assertCounter(t *testing.T, counter *Counter, value int) {
	t.Helper()
	if counter.Value() != value {
		t.Errorf("got %d, wanted %d", counter.Value(), value)
	}
}
