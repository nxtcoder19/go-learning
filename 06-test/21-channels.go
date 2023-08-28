package main

import (
	"fmt"
)

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	ch1 := make(chan string, 2)

	ch1 <- "buffer1"
	ch1 <- "buffer2"

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

}
