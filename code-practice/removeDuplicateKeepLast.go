// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"
	"strings"
)

func removeDuplicateKeepLast(s string) string {
	words := strings.Fields(s)
	seen := make(map[string]bool)
	result := []string{}

	for i := len(words) - 1; i >= 0; i-- {
		if !seen[words[i]] {
			seen[words[i]] = true
			result = append([]string{words[i]}, result...)
		}
	}
	return strings.Join(result, " ")
}

func main() {
	// fmt.Println("Try programiz.pro")
	input := "my name is piyush kumar i like people calling me by the name piyush my hobbies are  playing cricket my hobbies are  playing carrom"
	output := removeDuplicateKeepLast(input)
	fmt.Println(output)
}
