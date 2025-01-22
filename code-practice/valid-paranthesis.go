package main

import "fmt"

func canBeValid(s string, locked string) bool {
	n := len(s)

	// If the string length is odd, it's impossible to make it valid.
	if n%2 != 0 {
		return false
	}

	// Forward pass to ensure no unmatched ')'
	open := 0     // Tracks open '('
	flexible := 0 // Tracks flexible positions ('0' in locked)

	for i := 0; i < n; i++ {
		if locked[i] == '1' {
			if s[i] == '(' {
				open++
			} else { // s[i] == ')'
				if open > 0 {
					open--
				} else if flexible > 0 {
					flexible--
				} else {
					return false // Too many ')'
				}
			}
		} else {
			flexible++
		}
	}

	// Backward pass to ensure no unmatched '('
	close := 0
	flexible = 0

	for i := n - 1; i >= 0; i-- {
		if locked[i] == '1' {
			if s[i] == ')' {
				close++
			} else { // s[i] == '('
				if close > 0 {
					close--
				} else if flexible > 0 {
					flexible--
				} else {
					return false // Too many '('
				}
			}
		} else {
			flexible++
		}
	}

	return true
}

func main() {
	// Example 1
	s1 := "))()))"
	locked1 := "010100"
	fmt.Println(canBeValid(s1, locked1)) // Output: true

	// Example 2
	s2 := "()()"
	locked2 := "0000"
	fmt.Println(canBeValid(s2, locked2)) // Output: true

	// Example 3
	s3 := ")"
	locked3 := "0"
	fmt.Println(canBeValid(s3, locked3)) // Output: false
}

//A parentheses string is a non-empty string consisting only of '(' and ')'. It is valid if any of the following conditions is true:
//
//It is ().
//It can be written as AB (A concatenated with B), where A and B are valid parentheses strings.
//It can be written as (A), where A is a valid parentheses string.
//You are given a parentheses string s and a string locked, both of length n. locked is a binary string consisting only of '0's and '1's. For each index i of locked,
//
//If locked[i] is '1', you cannot change s[i].
//But if locked[i] is '0', you can change s[i] to either '(' or ')'.
//Return true if you can make s a valid parentheses string. Otherwise, return false.
