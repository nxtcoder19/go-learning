package main

import "fmt"

func getManipulatedArrayNew(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{} // Return empty slice if input is empty
	}

	result := make([]int, n)

	for i := range arr {
		left := 0  // Default to 0 if i-1 is out of bounds
		right := 0 // Default to 0 if i+1 is out of bounds

		if i > 0 {
			left = arr[i-1]
		}
		if i < n-1 {
			right = arr[i+1]
		}

		result[i] = left + arr[i] + right
	}

	return result
}

func getManipulatedArray(arr []int) []int {
	result := make([]int, len(arr))
	for i, _ := range arr {
		if i == 0 {
			result[i] = 0 + arr[i] + arr[i+1]
		} else if i == len(arr)-1 {
			result[i] = arr[i-1] + arr[i] + 0
		} else {
			result[i] = arr[i-1] + arr[i] + arr[i+1]
		}
	}
	return result
}

func main() {
	arr := []int{4, 0, 1, -2, 3}
	res := getManipulatedArray(arr)
	fmt.Println(res)
	resNew := getManipulatedArrayNew(arr)
	fmt.Println(resNew)
}
