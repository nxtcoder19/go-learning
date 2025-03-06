// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"
	// "math/rand"
	"sync"
	// "time"
)

type Employee1 struct {
	Name string
	Age  int
}

func generateEmployee1(count int) []interface{} {
	employees := make([]interface{}, count)

	for i := 0; i < count; i++ {
		employees[i] = Employee1{
			Name: fmt.Sprintf("Employee-%d", i+1),
			Age:  i + 1,
		}
		// time.Sleep(1 * time.Second)
	}
	return employees
}

func filterEmployees1(employees []interface{}, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	for _, emp := range employees {
		if e, ok := emp.(Employee1); ok && e.Age > 25 {
			ch <- e.Name
		}
	}

}

func main() {
	//   fmt.Println("Try programiz.pro")
	employees := generateEmployee1(100)

	ch := make(chan string, 10)
	// empCh := make(chan interface{}, 10)
	var wg sync.WaitGroup

	wg.Add(1)

	go filterEmployees1(employees, &wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for name := range ch {
		fmt.Println(name)
	}

}
