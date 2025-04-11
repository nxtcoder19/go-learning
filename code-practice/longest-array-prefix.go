package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Longest Common Prefix:", longestCommonPrefix(strs))
}
