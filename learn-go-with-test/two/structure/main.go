package structure

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Width float64
}

func Perimeter(rectangle Rectangle) float64{
	return 2 * (rectangle.Width + rectangle.Height)
}

func (c Circle) Area() float64 {
	return c.Width * c.Width * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}