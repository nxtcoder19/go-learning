package main

import (
	"fmt"
	"strings"
)

func matchedStrings(words []string) []string {
	var result []string
	for i, word := range words {
		for j, other := range words {
			if j != i && strings.Contains(other, word) {
				result = append(result, word)
				break
			}
		}
	}
	return result
}

func main() {
	s := []string{"mass", "as", "hero", "superhero", "ass"}
	fmt.Println(matchedStrings(s))

}
