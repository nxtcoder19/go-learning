package main

import "fmt"

func getSortedArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 11, 14, 6, 18, 8, 9, 10}
	fmt.Println(getSortedArray(arr))
}
