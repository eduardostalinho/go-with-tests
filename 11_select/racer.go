package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aTime := measureResponseTime(a)
	bTime := measureResponseTime(b)

	if aTime > bTime {
		return b
	}
	return a
}

func measureResponseTime(u string) time.Duration {
	start := time.Now()
	_, _ = http.Get(u)
	return time.Since(start)
}
