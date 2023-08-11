package main

import (
	"fmt"
)

type example interface {
	square() int
}

type sample int

func (s sample) square() int {
	// s = 30, effects only this function scope
	return int(s) * int(s)
}

type sample2 int

func (s *sample2) square() int {
	*s = 31
	return int(*s) * int(*s)
}

func squareExample(ex example) int {
	return ex.square()
}

func main() {
	// TODO: empty interface{}
	var x any = 5
	x = "abc"
	fmt.Printf("x = %s, type: %T\n", x, x)

	var n example
	n = sample(10)

	fmt.Printf("n = %d, type: %T\n", n, n)
	fmt.Println(n.square())
	fmt.Println(n)

	var m example
	z := sample2(13)
	m = &z

	fmt.Printf("m = %d, type: %T\n", m, m)
	fmt.Println(m.square())
	fmt.Println(z)

	fmt.Printf("square example: %d\n", squareExample(n))
	fmt.Printf("square example: %d\n", squareExample(m))
}
