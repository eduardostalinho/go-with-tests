package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 70
	hourHandLength   = 50
	clockCenter      = 150
)

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenter, p.Y + clockCenter}
	return p
}

func SecondHand(tm time.Time) Point {
	return makeHand(secondHandPoint(tm), secondHandLength)
}

func MinuteHand(tm time.Time) Point {
	return makeHand(minuteHandPoint(tm), minuteHandLength)
}

func HourHand(tm time.Time) Point {
	return makeHand(hourHandPoint(tm), hourHandLength)
}

func SVGWriter(w io.Writer, t time.Time) {
	sh := SecondHand(t)
	mh := MinuteHand(t)
	hh := HourHand(t)

	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	io.WriteString(w, HandLineTag(sh, "f00"))
	io.WriteString(w, HandLineTag(mh, "000"))
	io.WriteString(w, HandLineTag(hh, "000"))
	io.WriteString(w, svgEnd)
}

func HandLineTag(p Point, color string) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.2f" y2="%.2f" style="fill:none;stroke:#%s;stroke-width:3px;"/>`, p.X, p.Y, color)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
