package shapes

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(re Rectangle) float64 {
	return 2 * (re.Width + re.Height)
}

func (re Rectangle )Area() float64 {
	return re.Width * re.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}