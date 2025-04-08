package main

import "fmt"

// type Shape interface {
// 	Area() float64
// }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Area1() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {

	c := Circle{
		Radius: 5,
	}

	// var rad *Circle
	// rad = &c

	fmt.Println(c.Area())

	fmt.Println("Hello world")
}
