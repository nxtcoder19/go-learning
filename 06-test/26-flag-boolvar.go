package main

import (
	"flag"
	"fmt"
)

func main() {
	var verbose bool
	// Define a boolean flag named "verbose"
	flag.BoolVar(&verbose, "verbose", true, "Enable verbose mode")

	// Parse command-line arguments
	flag.Parse()

	// Use the value of the "verbose" flag
	if verbose {
		fmt.Println("Verbose mode is enabled.")
	} else {
		fmt.Println("Verbose mode is disabled.")
	}
}
