package main

import (
	"fmt"
)

func main() {
	// var x [2]int
	// x = [2]int{1, 2};

	// var x = [2]int{1, 2}

	x := [2]int{1, 2}

	fmt.Println(x)

	for i := range x {
		fmt.Printf("type1: index: %d, value: %d\n", i, x[i])
	}

	for i, v := range x {
		fmt.Printf("type2: index: %d, value: %d\n", i, v)
	}

	for _, v := range x {
		fmt.Printf("type3: value: %d\n", v)
	}

	// TODO: slices
}
