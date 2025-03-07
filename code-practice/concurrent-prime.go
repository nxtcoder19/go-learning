package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to check if a number is prime
func isPrimeNumber(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Producer function - Generates prime numbers less than 100 and sends them to a channel
func generatePrimes(primeCh chan int) {
	for i := 2; i < 100; i++ {
		if isPrimeNumber(i) {
			primeCh <- i
			time.Sleep(100 * time.Millisecond) // Simulating delay
		}
	}
	close(primeCh) // Close channel after sending all primes
}

// Consumer function - Reads primes and prints them
func printPrimes(primeCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for prime := range primeCh {
		fmt.Printf("Prime Number: %d\n", prime)
		time.Sleep(200 * time.Millisecond) // Simulating processing delay
	}
}

func main() {
	primeCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1) // Adding one task for the consumer

	// Start producer
	go generatePrimes(primeCh)

	// Start consumer
	go printPrimes(primeCh, &wg)

	wg.Wait() // Wait for consumer to finish processing
	fmt.Println("All tasks completed")
}
