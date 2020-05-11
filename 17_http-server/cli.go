package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please input the number of players: "

type CLI struct {
	game IGame
	in   *bufio.Scanner
	out  io.Writer
}

func NewCLI(game IGame, in io.Reader, out io.Writer) *CLI {
	return &CLI{game, bufio.NewScanner(in), out}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
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
