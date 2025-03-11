// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func generateFactorial(factCh chan int, limit int) {
// 	fact := 1
// 	for i := 1; i <= limit; i++ {
// 		fact *= i
// 		factCh <- fact
// 		time.Sleep(150 * time.Millisecond)
// 	}
// 	close(factCh)
// }

// func printFactorial(factCh chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for fact := range factCh {
// 		fmt.Printf("Factorial Number: %d\n", fact)
// 		time.Sleep(250 * time.Millisecond)
// 	}
// }

// func main() {
// 	factCh := make(chan int)
// 	limit := 10

// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go generateFactorial(factCh, limit)

// 	go printFactorial(factCh, &wg)

// 	wg.Wait() // Wait for consumers to finish processing
// 	fmt.Println("All Factorial numbers printed")
// }

package main

import (
	"fmt"
	"sync"
)

// Producer function - Generates Factorial numbers and sends them to a channel
func generateFactorial(factCh chan int, limit int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that the producer is done

	fact := 1
	for i := 1; i <= limit; i++ {
		fact *= i
		factCh <- fact
	}
	close(factCh) // Close channel after sending all Factorial numbers
}

// Consumer function - Reads Factorial numbers and prints them
func printFactorial(factCh chan int) {
	for fact := range factCh {
		fmt.Printf("Factorial Number: %d\n", fact)
	}
}

func main() {
	factCh := make(chan int)
	limit := 10

	var wg sync.WaitGroup
	wg.Add(1) // Adding one task for the producer

	// Start producer
	go generateFactorial(factCh, limit, &wg)

	// Start consumer (no need for WaitGroup in consumer, it will exit when channel is closed)
	printFactorial(factCh)

	wg.Wait() // Wait for producer to finish
	fmt.Println("All Factorial numbers printed")
}
