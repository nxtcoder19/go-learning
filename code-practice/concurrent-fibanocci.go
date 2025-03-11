package main

import (
	"fmt"
	"sync"
	"time"
)

// Producer function - Generates Fibonacci numbers up to a limit and sends them to a channel
func generateFibonacci(fibCh chan int, limit int) {
	a, b := 0, 1
	for a <= limit {
		fibCh <- a
		a, b = b, a+b
		time.Sleep(100 * time.Millisecond) // Simulating delay
	}
	close(fibCh) // Close channel after sending all Fibonacci numbers
}

// Consumer function - Reads Fibonacci numbers and prints them
func printFibonacci(fibCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for fib := range fibCh {
		fmt.Printf("Fibonacci Number: %d\n", fib)
		time.Sleep(200 * time.Millisecond) // Simulating processing delay
	}
}

func main() {
	fibCh := make(chan int) // Channel for Fibonacci numbers
	limit := 100            // Limit for Fibonacci numbers

	var wg sync.WaitGroup
	wg.Add(1) // Adding one task for the consumer

	// Start producer
	go generateFibonacci(fibCh, limit)

	// Start consumer
	go printFibonacci(fibCh, &wg)

	wg.Wait() // Wait for consumer to finish processing
	fmt.Println("All Fibonacci numbers printed")
}
