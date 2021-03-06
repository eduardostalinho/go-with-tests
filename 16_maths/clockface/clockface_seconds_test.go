package clockface

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	t.Run("at midnight", func(t *testing.T) {
		tm := simpleTime(0, 0, 0)
		want := Point{X: 150, Y: 150 - 90}

		got := SecondHand(tm)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("at 30 seconds", func(t *testing.T) {
		tm := simpleTime(0, 0, 30)
		want := Point{X: 150, Y: 150 + 90}

		got := SecondHand(tm)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("at %d seconds", test.time.Second()), func(t *testing.T) {
			got := secondHandPoint(test.time)
			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("got %v, want %v", got, test.point)
			}
		})
	}
}

func TestSecondInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), math.Pi / 2 * 3},
		{simpleTime(0, 0, 7), math.Pi / 30 * 7},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			want := test.angle

			got := secondInRadians(test.time)
			if !roughlyEqualFloat64(got, want) {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}
}
