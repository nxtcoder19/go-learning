package main

import "fmt"

func worker(c chan string) {
	c <- "Hello from Goroutine"
}

func main() {
	ch := make(chan string) // Create a channel

	go worker(ch)        // Start goroutine
	message := <-ch      // Receive data from channel
	fmt.Println(message) // Output: Hello from Goroutine
}
