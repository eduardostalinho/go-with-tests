package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlert(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

func (a BlindAlerterFunc) ScheduleAlert(duration time.Duration, amount int) {
	a(duration, amount)
}

func StdoutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind raised to %d\n", amount)
	})
}
