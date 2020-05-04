package main

import (
	"os"
	"time"

	"github.com/eduardostalinho/go-with-tests/16_maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}