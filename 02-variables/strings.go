package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	s = "hello world"

	fmt.Printf("my string: %s\n", s)
	fmt.Printf("my uppercase string: %s\n", strings.ToUpper(s))
	fmt.Printf("my lowercase string: %s\n", strings.ToLower(s))
	fmt.Printf("my titlecase string: %s\n", strings.ToTitle(s))

	// TODO: strings.HasPrefix()
	// TODO: strings.HasSuffix()
}
