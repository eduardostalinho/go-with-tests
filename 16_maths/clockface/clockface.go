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
	minuteHandLength = 70
	clockCentreX     = 150
	clockCentreY     = 150
)

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p

}

func SecondHand(tm time.Time) Point {
	return makeHand(secondHandPoint(tm), secondHandLength)
}

func MinuteHand(tm time.Time) Point {
	return makeHand(minuteHandPoint(tm), minuteHandLength)
}

func secondHandPoint(tm time.Time) Point {
	angle := secondInRadians(tm)
	return angleToPoint(angle)
}

func minuteHandPoint(tm time.Time) Point {
	angle := minuteInRadians(tm)
	return angleToPoint(angle)
}

func secondInRadians(tm time.Time) float64 {
	return math.Pi / (30 / float64(tm.Second()))
}

func minuteInRadians(tm time.Time) float64 {
	return (secondInRadians(tm) / 60) +
		(math.Pi / (30 / float64(tm.Minute())))
}

func angleToPoint(angle float64) Point {
	return Point{math.Sin(angle), math.Cos(angle)}
}
