package clockface

import (
	"math"
	"time"
)

const (
	halfClockSeconds = 30
	clockSeconds     = halfClockSeconds * 2
	halfClockMinutes = 30
	clockMinutes     = halfClockMinutes * 2
	halfClockHours   = 6
	clockHours       = 12
)

type Point struct {
	X float64
	Y float64
}

func secondHandPoint(tm time.Time) Point {
	angle := secondInRadians(tm)
	return angleToPoint(angle)
}

func minuteHandPoint(tm time.Time) Point {
	angle := minuteInRadians(tm)
	return angleToPoint(angle)
}

func hourHandPoint(tm time.Time) Point {
	angle := hourInRadians(tm)
	return angleToPoint(angle)
}

func secondInRadians(tm time.Time) float64 {
	return math.Pi / (halfClockSeconds / float64(tm.Second()))
}

func minuteInRadians(tm time.Time) float64 {
	return (secondInRadians(tm) / clockMinutes) +
		(math.Pi / (halfClockMinutes / float64(tm.Minute())))
}

func hourInRadians(tm time.Time) float64 {
	return (minuteInRadians(tm) / clockHours) +
		(math.Pi / (halfClockHours / float64(tm.Hour()%12)))
}

func angleToPoint(angle float64) Point {
	return Point{math.Sin(angle), math.Cos(angle)}
}
