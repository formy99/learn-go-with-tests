package structs

import "math"

func Perimeter(r Rectangle) float64 {
	return (r.Height + r.Width) * 2
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
