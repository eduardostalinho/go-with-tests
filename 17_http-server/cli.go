package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const playerPrompt = "Please input the number of players: "

type CLI struct {
	game *Game
	in   *bufio.Scanner
	out  io.Writer
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	game := &Game{store, alerter}
	return &CLI{game, bufio.NewScanner(in), out}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, playerPrompt)
	numberOfPlayersInput := c.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	c.game.Start(numberOfPlayers)

	winnerInput := c.readLine()
	winner := extractWinner(winnerInput)

	c.game.Finish(winner)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
