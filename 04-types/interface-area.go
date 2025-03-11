package main

import (
	"fmt"
)

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

type Rectangle struct {
	length, width float64
}

type Triangle struct {
	base, height float64
}

func (s *Square) area() float64 {
	return s.side * s.side
}

func (r *Rectangle) area() float64 {
	return r.length * r.width
}

func (t *Triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func newSquare(side float64) Shape {
	return &Square{side: side}
}

func newRectangle(length, width float64) Shape {
	return &Rectangle{length: length, width: width}
}

func newTriangle(base, height float64) Shape {
	return &Triangle{base: base, height: height}
}

func main() {
	fmt.Println("Shape Area Calculation")

	sq := newSquare(5)
	rect := newRectangle(10, 4)
	tri := newTriangle(6, 3)

	fmt.Println("Square Area:", sq.area())
	fmt.Println("Rectangle Area:", rect.area())
	fmt.Println("Triangle Area:", tri.area())
}
