package method

import "math"

type Triangle struct {
	Base   float64
	Height float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return t.Height * t.Base * 0.5
}
