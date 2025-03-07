package main

import (
	"fmt"
	"sync"
	"time"
)

// Producer function - Generates numbers from 1 to 100 and sends them to the channel
func generateNumbers(numCh chan int) {
	for i := 1; i <= 10; i++ {
		numCh <- i
		time.Sleep(50 * time.Millisecond) // Simulating delay
	}
	close(numCh) // Close channel after sending all numbers
}

// Consumer function - Prints even numbers
func printEvenNumbers(numCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numCh {
		if num%2 == 0 {
			fmt.Printf("Even Number: %d\n", num)
			time.Sleep(100 * time.Millisecond) // Simulating processing delay
		}
	}
}

// Consumer function - Prints odd numbers
func printOddNumbers(numCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numCh {
		if num%2 != 0 {
			fmt.Printf("Odd Number: %d\n", num)
			time.Sleep(100 * time.Millisecond) // Simulating processing delay
		}
	}
}

func main() {
	numCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2) // Adding two tasks for the consumers

	// Start producer
	go generateNumbers(numCh)

	// Start consumers
	go printEvenNumbers(numCh, &wg)
	go printOddNumbers(numCh, &wg)

	wg.Wait() // Wait for consumers to finish processing
	fmt.Println("All tasks completed")
}
