// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	ch := make(chan int, 3)
	go square(2, ch, &wg)
	go square(3, ch, &wg)
	go square(4, ch, &wg)
	wg.Wait()
	close(ch)
	for result := range ch {
		fmt.Println("Result:", result)
	}
}

func square(num int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 	mu.lock()
	ch <- num * num
	// mu.unLock()
}
