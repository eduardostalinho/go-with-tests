package shapes

import "math"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base float64
	Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func Perimeter(s Shape) float64 {
	return s.Perimeter()
}

func Area(s Shape) float64 {
	return s.Area()
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height

}

func (c Circle) Perimeter() float64 {
	return (2 * c.Radius * math.Pi) * (2 * c.Radius * math.Pi)
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}



func (t Triangle) Area() float64 {
	return  (t.Base * t.Height) / 2
}

func (t Triangle) Perimeter() float64 {
	return 0
}