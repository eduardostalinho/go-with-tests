package clockface

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestMinuteHand(t *testing.T) {
	t.Run("at midnight", func(t *testing.T) {
		tm := simpleTime(0, 0, 0)
		want := Point{X: 150, Y: 150 - 70}

		got := MinuteHand(tm)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("at 30 minutes", func(t *testing.T) {
		tm := simpleTime(0, 30, 0)
		want := Point{X: 150, Y: 150 + 70}

		got := MinuteHand(tm)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

}
func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("at %d minutes", test.time.Minute()), func(t *testing.T) {
			got := minuteHandPoint(test.time)
			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("got %v, want %v", got, test.point)
			}
		})
	}
}

func TestMinuteInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			want := test.angle

			got := minuteInRadians(test.time)
			if !roughlyEqualFloat64(got, want) {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}
}
