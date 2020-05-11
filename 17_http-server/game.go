package poker

import (
	"time"
)

type Game struct {
	store   PlayerStore
	alerter BlindAlerter
}

type IGame interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

func NewGame(store PlayerStore, alerter BlindAlerter) *Game {
	return &Game{store, alerter}
}

func (g *Game) Start(numberOfPlayers int) {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	timeIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	for _, blind := range blinds {
		g.alerter.ScheduleAlert(blindTime, blind)
		blindTime = blindTime + timeIncrement
	}
}

func (g *Game) Finish(winner string) {
	g.store.RecordWin(winner)
}
