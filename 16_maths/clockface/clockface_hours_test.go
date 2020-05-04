package clockface

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestHourHand(t *testing.T) {
	t.Run("at midnight", func(t *testing.T) {
		tm := simpleTime(0, 0, 0)
		want := Point{X: 150, Y: 150 - 50}

		got := HourHand(tm)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("at 6 hours", func(t *testing.T) {
		tm := simpleTime(6, 0, 0)
		want := Point{X: 150, Y: 150 + 50}

		got := HourHand(tm)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

}
func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(9, 0, 0), Point{-1, 0}},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("at %d hours", test.time.Hour()), func(t *testing.T) {
			got := hourHandPoint(test.time)
			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("got %v, want %v", got, test.point)
			}
		})
	}
}

func TestHourInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			want := test.angle

			got := hourInRadians(test.time)
			if !roughlyEqualFloat64(got, want) {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}
}
