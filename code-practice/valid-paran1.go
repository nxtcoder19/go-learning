package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	matchingBrackets := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, ch := range s {
		// If closing bracket, check if stack has the matching opening bracket
		if openBracket, exists := matchingBrackets[ch]; exists {
			//fmt.Printf("Opening bracket: %c\n", openBracket)
			//fmt.Printf("exists bracket: %c\n", exists)
			//fmt.Printf("if stack: %c\n", ch)
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			stack = stack[:len(stack)-1] // Pop from stack
		} else {
			// If opening bracket, push onto stack
			//fmt.Printf("Pushing onto stack: %c\n", ch)
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0 // Valid if stack is empty
}

func main() {
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"",
	}

	for _, test := range testCases {
		fmt.Printf("Input: %s -> Valid: %v\n", test, isValid(test))
	}
}
