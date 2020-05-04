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
	Circle  struct {
		Text  string `xml:",chardata"`
		Cx    string `xml:"cx,attr"`
		Cy    string `xml:"cy,attr"`
		R     string `xml:"r,attr"`
		Style string `xml:"style,attr"`
	} `xml:"circle"`
	Line struct {
		Text  string `xml:",chardata"`
		X1    string `xml:"x1,attr"`
		Y1    string `xml:"y1,attr"`
		X2    string `xml:"x2,attr"`
		Y2    string `xml:"y2,attr"`
		Style string `xml:"style,attr"`
	} `xml:"line"`
}

func TestSVGWriterAtMidnight(t *testing.T) {
	b := bytes.Buffer{}
	tm := simpleTime(0, 0, 0)

	SVGWriter(&b, tm)

	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := "150"
	y2 := "60"

	if !(svg.Line.X2 == x2) || !(svg.Line.Y2 == y2) {
		t.Errorf("Expected line x2 %s at and y2 at %s, got SVG %s", x2, y2, b.String())
	}

}
