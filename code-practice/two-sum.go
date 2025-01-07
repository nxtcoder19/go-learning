package main

import "fmt"

func twoSum(arr []int, target int) []int {
	firstIndex := 0
	secondIndex := 1
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				firstIndex = i
				secondIndex = j
			}
		}
	}
	return []int{firstIndex, secondIndex}
}

func main() {
	var arr []int
	arr = []int{2, 11, 15, 7}
	fmt.Println(twoSum(arr, 9))
}
