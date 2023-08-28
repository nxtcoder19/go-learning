package main

import (
	"flag"
	"fmt"
)

func main() {

	var name string
	var age int

	flag.StringVar(&name, "name", "John Doe", "Your name")
	flag.IntVar(&age, "age", 0, "Your age")

	// Parse the command line flags
	// flag.Parse()

	// Print the values of the flags
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)

	// var word string

	wordPtr := flag.String("word", "word", "foo")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// Execurion process:

// go build abc.go
// abc command line will be created
// ./abc -name=aaa -age=12 -word=opt a1 a2 a3
