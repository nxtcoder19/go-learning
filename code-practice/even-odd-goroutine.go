package main

import (
	"fmt"
	"sync"
)

func printOdd(oddChan chan bool, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		<-oddChan // Wait for signal to print odd number
		fmt.Println("Odd:", i)
		evenChan <- true // Signal even goroutine
	}
}

func printEven(evenChan chan bool, oddChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-evenChan // Wait for signal to print even number
		fmt.Println("Even:", i)
		oddChan <- true // Signal odd goroutine
	}
}

func main() {
	oddChan := make(chan bool, 1)  // Buffered channel to prevent blocking
	evenChan := make(chan bool, 1) // Buffered channel to prevent blocking
	var wg sync.WaitGroup

	wg.Add(2)
	go printOdd(oddChan, evenChan, &wg)
	go printEven(evenChan, oddChan, &wg)

	oddChan <- true // Start with odd
	wg.Wait()
	close(oddChan)
	close(evenChan)
}
