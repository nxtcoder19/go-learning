package main

import (
	"fmt"
	"regexp"
	"strings"
)

//func reverseString(s []string) []string {
//	//result := make([]string, 0)
//	//result := make([]string, len(s))
//	var result []string
//	for i := len(s) - 1; i >= 0; i-- {
//		result = append(result, s[i])
//	}
//
//	return result
//}

func removeExtraSpaces(input string) string {
	// Create a regex pattern to match two or more spaces
	re := regexp.MustCompile(`\s{2,}`)
	// Replace matches with a single space
	result := re.ReplaceAllString(input, " ")
	// Trim leading and trailing spaces
	return strings.TrimSpace(result)
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	fmt.Println(words)
	//var reversedWords string
	var reverseArrayStrings []string
	for i := len(words) - 1; i >= 0; i-- {
		//reversedWords += words[i] + " "
		reverseArrayStrings = append(reverseArrayStrings, words[i])
	}
	fmt.Println(strings.Join(reverseArrayStrings, " "))
	result := strings.Join(reverseArrayStrings, " ")
	return removeExtraSpaces(result)
}

func main() {
	//var s []string
	//s = []string{"the", "sky", "is", "blue"}
	//fmt.Println(reverseString(s))
	fmt.Println(reverseWords("  hello world  "))
}
