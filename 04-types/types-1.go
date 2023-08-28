package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (mf MyFloat) Abs() float64 {
	if mf < 0 {
		return float64(-mf)
	}
	return float64(mf)
}

type vertex struct {
	X, Y float64
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type square struct {
	a, b int
}

func (s *square) Abs() int {
	return s.a * s.b
}

func main() {
	var a Abser
	a = MyFloat(-5.0)
	fmt.Println("abs %d", a.Abs())

	d := vertex{2, 3}
	fmt.Println("vertex", d.Abs())

	e := square{3, 4}
	fmt.Println(e.Abs())

}
