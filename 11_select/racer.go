package racer

import (
	"net/http"
	"time"
)

const (
	ErrTimeoutErr    = RacerError("Timed out waiting for response")
	tenSecondTimeout = 10 * time.Second
)

type RacerError string

func (e RacerError) Error() string {
	return string(e)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", ErrTimeoutErr
	}

}
func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ping(u string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(u)
		close(ch)
	}()
	return ch
}
