package clockface

import (
	"fmt"
	"io"
	"time"
)

func SVGWriter(w io.Writer, t time.Time) {
	sh := SecondHand(t)
	mh := MinuteHand(t)

	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	io.WriteString(w, HandLineTag(sh))
	io.WriteString(w, HandLineTag(mh))
	io.WriteString(w, svgEnd)
}

func HandLineTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.f" y2="%.f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
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
