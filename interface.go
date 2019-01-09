package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{Radius: 7.0}
	rectangle := Rectangle{Width: 10.0, Height: 10.0}

	fmt.Printf("Area of Circle = %f\n", getArea(circle))
	fmt.Printf("Area of Rectangle = %f\n", getArea(rectangle))
}

// Polymorphism
// Generic
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func getArea(s Shape) float64 {
	return s.Area()
}
