package main

import (
	"fmt"
)

// Function to check if a character is a vowel
func isVowel(ch byte) bool {
	vowels := "aeiouy"
	for i := range vowels {
		//fmt.Println(i)
		//fmt.Println(vowels[i])

		if ch == vowels[i] {
			return true
		}
	}
	return false
}

// Function to count valid substrings matching the pattern
func countMatchingSubstrings(pattern, source string) int {
	pLen := len(pattern)
	sLen := len(source)
	count := 0

	// Iterate over all possible substrings of source with length equal to pattern
	for i := 0; i <= sLen-pLen; i++ {
		match := true
		for j := 0; j < pLen; j++ {
			if (pattern[j] == '0' && !isVowel(source[i+j])) || (pattern[j] == '1' && isVowel(source[i+j])) {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	return count
}

func main() {
	pattern := "010"
	source := "amazing"

	result := countMatchingSubstrings(pattern, source)
	fmt.Println("Number of matching substrings:", result)
}
