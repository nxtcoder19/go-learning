package main

import (
	"fmt"
	"strings"
)

func getPrefixCount(words []string, prefix string) int {
	prefixCount := 0
	for _, word := range words {
		if strings.HasPrefix(word, prefix) {
			prefixCount++
		}
	}
	return prefixCount
}

func main() {
	var s []string
	s = []string{"pay", "attention", "practice", "attend"}
	pref := "at"
	fmt.Println(getPrefixCount(s, pref))

}
