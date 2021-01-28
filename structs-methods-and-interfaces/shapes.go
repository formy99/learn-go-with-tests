package structs_methods_and_interfaces

import "math"

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

type Triangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

//func (s Shape) Area() float64 {
//	return 0
//}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Area() float64 {
	return t.Height * t.Width / 2
}
