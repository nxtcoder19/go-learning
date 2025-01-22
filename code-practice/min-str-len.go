package main

import (
	"fmt"
)

func getMinLength(s string) int {
	m := map[rune]int{}
	for _, r := range s {
		fmt.Println(r)
		m[r]++
		fmt.Println(m)
	}

	length := 0
	for _, v := range m {
		if v%2 == 1 {
			length++
		} else {
			length += 2
		}
	}

	return length
}

func main() {
	s := "abaacbcbb"
	fmt.Println(getMinLength(s))
}

//You are given a string s.
//
//You can perform the following process on s any number of times:
//
//Choose an index i in the string such that there is at least one character to the left of index i that is equal to s[i], and at least one character to the right that is also equal to s[i].
//Delete the closest character to the left of index i that is equal to s[i].
//Delete the closest character to the right of index i that is equal to s[i].
//Return the minimum length of the final string s that you can achieve.
