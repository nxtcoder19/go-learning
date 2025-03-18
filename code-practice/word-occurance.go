package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "india is my country, i love my country"

	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, ",", "")

	words := strings.Fields(text)
	fmt.Println(words)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		fmt.Println(word, "-", count)
	}
}
