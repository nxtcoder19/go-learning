package main

import (
	"fmt"
	"time"
)

// / buffered channel without consumer
func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i // This will block when the buffer is full
	}
	fmt.Println("Producer finished producing")
}

func main() {
	bufferSize := 5
	ch := make(chan int, bufferSize) // Buffered channel with capacity 5

	go produce(ch)

	// No consumer, so the producer will block when the buffer is full
	time.Sleep(3 * time.Second) // Simulating delay, eventually causing deadlock

	fmt.Println("Main function completed") // This may not execute if producer is blocked
}

/// unbuffered channel without consumer

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func produce(ch chan int) {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Printf("Producing: %d\n", i)
// 		ch <- i // This will block indefinitely if no consumer is present
// 	}
// 	fmt.Println("Producer finished producing") // This will never be reached
// }

// func main() {
// 	ch := make(chan int) // Unbuffered channel

// 	go produce(ch) // Start producer

// 	// No consumer, so producer will block indefinitely
// 	// select {} // Block main function indefinitely (deadlock)

// 	time.Sleep(4 * time.Second)
// 	fmt.Println("All tasks completed")
// }
