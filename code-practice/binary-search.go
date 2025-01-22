package main

import "fmt"

func binarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
	found := false
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			found = true
		}
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return found
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := binarySearch(arr, 2)
	fmt.Println(x)
}
