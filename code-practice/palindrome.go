package main

import "fmt"

func isPalindrome(s int) bool {
	res := fmt.Sprintf("%d", s)
	reversed := ""
	for i := len(res) - 1; i >= 0; i-- {
		reversed += string(res[i])
	}
	fmt.Println(res, reversed)
	return res == reversed
}

func main() {
	s := 10
	fmt.Println(isPalindrome(s))
}
