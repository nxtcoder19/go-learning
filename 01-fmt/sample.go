package main

import (
	"fmt"
	"os"
)

func printf(msg string, args ...any) {
	fmt.Fprintf(os.Stdout, msg, args...)
}

func main() {
	// 0 => stdin
	// 1=> stdout
	// 2 => stderr
	fmt.Println("Hello world")
	fmt.Printf("hello %s, greetings from %s\n", "piyush", "golang")

	fmt.Fprintf(os.Stdout, "hello not from this world\n")
	printf("hello %s", "mohit")

	fmt.Fprintf(os.Stderr, "hello this is error\n")
}
