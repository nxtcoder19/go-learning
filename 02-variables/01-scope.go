package main

import (
	"fmt"
)

func main() {
	// var a int = 5
	var a int
	a = 5

	{
		a := true
		fmt.Println(a)
	}

	if true {
		// a := "asdaf"
		fmt.Println(a)
	}
	// a = "Asdfasf"

	fmt.Println(a)

	// TODO: float32()
	var b float32
	b = 43534453345
	fmt.Println(b)

	// TODO: float64()
}
