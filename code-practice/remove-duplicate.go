package main

import "fmt"

func removeDuplicates(arr []int) []int {
	var newArr []int
	arrMap := make(map[int]bool)

	for _, v := range arr {
		if _, ok := arrMap[v]; !ok {
			arrMap[v] = true
			newArr = append(newArr, v)
		}
	}ww
	return newArr
}

func main() {
	arr := []int{1, 2, 3, 4, 10, 9, 6, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(removeDuplicates(arr))
}
