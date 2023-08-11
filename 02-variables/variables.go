package main

import (
	"fmt"
)

func main() {
	age := int8(127)
	// int32(5)
	// int8
	fmt.Printf("age: %d, type: %T\n", age, age)

	s := "asdfasdfaf"

	length := uint(len(s))
	fmt.Printf("string: %s, length: %d %T\n", s, len(s), len(s))

	if length > 0 {
		fmt.Println("here")
	}
}
