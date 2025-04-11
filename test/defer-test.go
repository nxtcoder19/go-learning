package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	func() {
		defer fmt.Println("Deferred inside function")
		numbers = numbers[:len(numbers)-1]
	}()

	fmt.Println("After modification:", numbers)
}
