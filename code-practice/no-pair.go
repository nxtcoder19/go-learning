package main

import (
	"fmt"
)

func checkOddPair(arr []string) []int {
	res := make([]int, 0)
	for _, s := range arr {
		res = append(res, getOddCount(s))
	}
	return res
}

func getOddCount(s string) int {
	count := 0
	for i := 0; i < len(s)-1; i++ { // Use len(s)-1 to avoid index out of range
		if s[i] == s[i+1] {
			count++
			i++
		}
	}
	return count
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(checkOddPair([]string{"abb", "def", "hhi", "jkl", "aab", "abaaabaaab", "abaaabaaabccdgdde"}))
}
