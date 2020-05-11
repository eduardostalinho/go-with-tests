package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	store   PlayerStore
	in      *bufio.Scanner
	alerter BlindAlerter
}

type BlindAlerter interface {
	ScheduleAlert(duration time.Duration, amount int)
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{store, bufio.NewScanner(in), alerter}
}

func (c *CLI) PlayPoker() {
	c.alerter.ScheduleAlert(1*time.Second, 1)
	line := c.readLine()
	winner := extractWinner(line)
	c.store.RecordWin(winner)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
