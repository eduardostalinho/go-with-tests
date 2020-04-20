package racer

import (
	"net/http"
)

func Racer(a, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
