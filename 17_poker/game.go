package poker

import (
	"io"
	"time"
)

type Game interface {
	Start(numberOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}

type TexasHoldEm struct {
	store   PlayerStore
	alerter BlindAlerter
}

func NewTexasHoldEm(store PlayerStore, alerter BlindAlerter) *TexasHoldEm {
	return &TexasHoldEm{store, alerter}
}

func (g *TexasHoldEm) Start(numberOfPlayers int, alertsDestination io.Writer) {

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	timeIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	for _, blind := range blinds {
		g.alerter.ScheduleAlert(blindTime, blind, alertsDestination)
		blindTime = blindTime + timeIncrement
	}
}

func (g *TexasHoldEm) Finish(winner string) {
	g.store.RecordWin(winner)
}
