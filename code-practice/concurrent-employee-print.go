package main

import (
	"fmt"
	"sync"
	"time"
)

// Employee struct
type Employee struct {
	Name string
	Age  int
}

// Producer function - Generates employees and sends them to a channel
func generateEmployeeNew(count int, empCh chan interface{}) {
	for i := 0; i < count; i++ {
		empCh <- Employee{
			Name: fmt.Sprintf("Employee-%d", i+1),
			Age:  i + 1,
		}
		time.Sleep(100 * time.Millisecond) // Simulating delay
	}
	close(empCh) // Close channel after sending all employees
}

// Consumer function - Reads employees and prints them
func printEmployeeNew(empCh chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for emp := range empCh {
		if e, ok := emp.(Employee); ok && e.Age > 25 {
			fmt.Printf("Consuming %s\n", e.Name)
			time.Sleep(200 * time.Millisecond) // Simulating processing delay
		}
	}
}

func main() {
	empCh := make(chan interface{})

	var wg sync.WaitGroup
	wg.Add(1) // Adding one task for the consumer

	// Start producer
	go generateEmployeeNew(100, empCh)

	// Start consumer
	go printEmployeeNew(empCh, &wg)

	wg.Wait() // Wait for consumer to finish processing
	fmt.Println("All tasks completed")
}
