package clockface

import (
	"bytes"
	"encoding/xml"
	"testing"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx    float64 `xml:"cx,attr"`
	Cy    float64 `xml:"cy,attr"`
	R     float64 `xml:"r,attr"`
}

type Line struct {
	X1    float64 `xml:"x1,attr"`
	Y1    float64 `xml:"y1,attr"`
	X2    float64 `xml:"x2,attr"`
	Y2    float64 `xml:"y2,attr"`
}

func TestSVGWriterAtMidnight(t *testing.T) {
	b := bytes.Buffer{}
	tm := simpleTime(0, 0, 0)

	SVGWriter(&b, tm)

	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := 150.
	y2 := 60.

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
	}
	t.Errorf("Expected line x2 %f at and y2 at %f, got SVG %s", x2, y2, b.String())
}
