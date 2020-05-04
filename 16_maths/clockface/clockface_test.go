package clockface

import (
	"fmt"
	"math"
	"time"
)

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(312, time.February, 5, h, m, s, 0, time.UTC)

}

func testName(t time.Time) string {
	return fmt.Sprintf("%d:%d:%d", t.Hour(), t.Minute(), t.Second())
}
