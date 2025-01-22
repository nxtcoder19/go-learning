package main

import (
	"fmt"
)

func producer(ch chan int, done chan bool) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i // Send data to the channel
	}
	close(ch)    // Close the channel when done
	done <- true // Notify completion
}

func consumer(ch chan int, done chan bool) {
	for value := range ch { // Receive data from the channel
		fmt.Printf("Consuming %d\n", value)
	}
	done <- true // Notify completion
}

func main() {
	ch := make(chan int)    // Channel for integers
	done := make(chan bool) // Channel to signal completion

	go producer(ch, done) // Start producer goroutine
	go consumer(ch, done) // Start consumer goroutine

	// Wait for both producer and consumer to complete
	<-done
	<-done

	fmt.Println("All tasks completed")
}
