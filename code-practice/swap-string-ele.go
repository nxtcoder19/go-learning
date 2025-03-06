package main

import (
	"fmt"
)

// Function to check if s1 can be transformed into s2 by swapping characters
func canTransform(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Frequency maps
	count1 := make(map[rune]int)
	count2 := make(map[rune]int)

	// Count character frequencies
	for _, ch := range s1 {
		count1[ch]++
	}
	for _, ch := range s2 {
		count2[ch]++
	}

	// Compare frequency maps manually
	for key, val := range count1 {
		if count2[key] != val {
			return false
		}
	}

	//fmt.Println(count1)
	//fmt.Println(count2)

	return true
}

func main() {
	s1 := "abcd"
	s2 := "bcda"

	if canTransform(s1, s2) {
		fmt.Println("Yes, s1 can be transformed into s2")
	} else {
		fmt.Println("No, s1 cannot be transformed into s2")
	}
}
