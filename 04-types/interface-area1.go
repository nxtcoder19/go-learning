package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
}

// Implement the interface for a struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Another struct implementing Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	printArea(c) // Circle satisfies Shape
	printArea(r) // Rectangle satisfies Shape
}
