package main

import "fmt"

func longestSubStr(s string) string {
	charMap := make(map[rune]int)
	maxLen := 0
	start := 0
	subStringStart := 0

	for i, ch := range s {
		if lastIndex, found := charMap[ch]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		charMap[ch] = i

		if i-start+1 > maxLen {
			maxLen = i - start + 1
			subStringStart = start
		}
	}

	return s[subStringStart : subStringStart+maxLen]

}

func main() {

	input := "abcabcbb"
	fmt.Println(longestSubStr(input))

	// fmt.Println("Try programiz.pro")
}
