package racer

import (
	"net/http"
	"time"
)

const ErrTimeoutErr = RacerError("Timed out waiting for response")

type RacerError string

func (e RacerError) Error() string {
	return string(e)
}

func Racer(a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", ErrTimeoutErr
	}
}

func ping(u string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(u)
		close(ch)
	}()
	return ch
}
