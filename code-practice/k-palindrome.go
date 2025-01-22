package main

import (
	"fmt"
)

func canConstruct(s string, k int) bool {
	// If k is greater than the length of the string, it's impossible
	if k > len(s) {
		return false
	}

	// Count character frequencies
	charFrequency := make(map[rune]int)
	for _, char := range s {
		charFrequency[char]++
	}

	// Count characters with odd frequencies
	oddCount := 0
	for _, count := range charFrequency {
		if count%2 != 0 {
			oddCount++
		}
	}

	// If the number of odd frequency characters exceeds k, return false
	return oddCount <= k
}

func main() {
	// Test cases
	fmt.Println(canConstruct("annabelle", 2)) // Output: true
	fmt.Println(canConstruct("leetcode", 3))  // Output: false
	fmt.Println(canConstruct("true", 4))      // Output: true
}
