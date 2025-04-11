package main

import (
	"fmt"
)

func findMissingElements(arr []int, start int, end int) []int {
	exists := make(map[int]bool)
	missing := []int{}

	for _, num := range arr {
		exists[num] = true
	}

	for i := start; i <= end; i++ {
		if !exists[i] {
			missing = append(missing, i)
		}
	}

	return missing
}

func main() {
	arr := []int{1, 3, 5, 6}
	start := 1
	end := 6
	fmt.Println("Missing elements:", findMissingElements(arr, start, end))
}
