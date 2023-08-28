package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func calc(x int, y int) (int, int, int, float64) {
	return x + y, x - y, x * y, float64(x) / float64(y)
}

func calc2(x int, y int) (sum int, sub int, mul int, div float64) {
	sum = x + y
	sub = x - y
	mul = x * y
	div = float64(x) / float64(y)

	return sum, sub, mul, div
}

func addTowNo() {
	fmt.Println("Add + double two numbers: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2))
}

func main() {
	addTowNo()
	fmt.Println(add(5, 6))
	fmt.Println(calc(5, 6))
	fmt.Println(calc2(9, 6))

	x := 5
	var y func() int
	y = func() int {
		return x
	}
	x = 7
	fmt.Printf("y: %T value: %d \n", y, y())
	x = 8
}
