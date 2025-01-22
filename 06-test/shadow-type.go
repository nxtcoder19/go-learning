package main

import "fmt"

func main() {
	x := 10 // Outer variable
	fmt.Println("Outer x:", x)

	{
		x := 20 // Shadows the outer variable
		fmt.Println("Inner x:", x)
	}

	fmt.Println("Outer x after block:", x) // Original x remains unaffected
}
