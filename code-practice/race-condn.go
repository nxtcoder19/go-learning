package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mu1 sync.Mutex

// func increment() {
// 	counter++
// }

func increment() {
	mu1.Lock()
	counter++
	mu1.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	time.Sleep(time.Second)
	fmt.Println("Final Counter:", counter)
}
