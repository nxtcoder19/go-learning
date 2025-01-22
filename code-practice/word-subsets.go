package main

import (
	"fmt"
	"strings"
)

func getSubsets(words1 []string, words2 []string) []string {
	var result []string

	// Helper function to count character frequency
	countChars := func(word string) map[rune]int {
		freq := make(map[rune]int)
		for _, ch := range word {
			freq[ch]++
		}
		return freq
	}

	// Create a maximum frequency map for words2
	maxFreq := make(map[rune]int)
	for _, word := range words2 {
		for ch, count := range countChars(word) {
			if maxFreq[ch] < count {
				maxFreq[ch] = count
			}
		}
	}

	// Check if a word from words1 contains all characters in the required frequency
	isSubset := func(word1 string) bool {
		wordFreq := countChars(word1)
		for ch, count := range maxFreq {
			if wordFreq[ch] < count {
				return false
			}
		}
		return true
	}

	// Filter words1
	for _, word1 := range words1 {
		if isSubset(word1) {
			result = append(result, word1)
		}
	}

	return result
}

func getWordSubsets(words1 []string, words2 []string) []string {
	var result []string
	isSubset := func(word1 string) bool {
		for _, word2 := range words2 {
			//if !strings.Contains(word1, word2) {
			//	return false
			//}
			for i := 0; i < len(word2); i++ {
				if !strings.Contains(word1, string(word2[i])) {
					return false
				}
			}
		}
		return true
	}

	for _, word1 := range words1 {
		if isSubset(word1) {
			result = append(result, word1)
		}
	}
	return result
}

func main() {
	var words1 []string
	words1 = []string{"amazon", "apple", "facebook", "google", "leetcode"}

	words2 := []string{"e", "oo"}
	fmt.Println(getWordSubsets(words1, words2))
	fmt.Println(getSubsets(words1, words2))
}
