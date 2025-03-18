package main

import (
	"fmt"
	"time"
)

func producer1(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i // Send data to the channel
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // Close the channel when done
}

// func consumer1(ch chan int) {
// 	for value := range ch { // Receive data from the channel
// 		fmt.Printf("Consuming %d\n", value)
// 		time.Sleep(700 * time.Millisecond)
// 	}
// }

func main() {
	ch := make(chan int) // Create a channel for integers

	go producer1(ch) // Start the producer goroutine
	// go consumer1(ch) // Start the consumer goroutine

	// Wait for the goroutines to finish
	time.Sleep(4 * time.Second)
	fmt.Println("All tasks completed")
}
