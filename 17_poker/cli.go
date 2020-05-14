package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	PlayerPrompt      = "Please input the number of players: "
	BadPlayerInputMsg = "Invalid input for number of players, please input a valid number."
	BadWinMsg         = "Invalid message for win. Please input `{Name} wins`"
)

type CLI struct {
	game Game
	in   *bufio.Scanner
	out  io.Writer
}

func NewCLI(game Game, in io.Reader, out io.Writer) *CLI {
	return &CLI{game, bufio.NewScanner(in), out}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numberOfPlayersInput := c.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	if err != nil {
		fmt.Fprintf(c.out, BadPlayerInputMsg)
		return
	}
	c.game.Start(numberOfPlayers, c.out)

	winnerInput := c.readLine()
	winner, err := extractWinner(winnerInput)
	if err != nil {
		fmt.Fprintf(c.out, BadWinMsg)
		return
	}

	c.game.Finish(winner)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(input string) (string, error) {
	suffix := " wins"
	if input[len(input)-len(suffix):] != suffix {
		return "", fmt.Errorf(BadWinMsg)
	}
	return strings.Replace(input, suffix, "", 1), nil
}
