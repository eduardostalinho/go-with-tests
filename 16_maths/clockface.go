package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

func SecondHand(tm time.Time) Point {
	p := secondHandPoint(tm)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}

func secondHandPoint(tm time.Time) Point {
	angle := secondInRadians(tm)
	return Point{math.Sin(angle), math.Cos(angle)}
}

func secondInRadians(tm time.Time) float64 {
	return math.Pi / (30 / float64(tm.Second()))
}
