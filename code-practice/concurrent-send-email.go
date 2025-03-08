// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"
	"sync"
	"time"
)

func generateEmailWithMessage(msgCh chan string, emailList []string, msg string) {
	for i := 0; i < len(emailList); i++ {
		msgCh <- emailList[i] + msg
		time.Sleep(50 * time.Millisecond)
	}
	close(msgCh)
}

func printSendMessage(msgCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range msgCh {
		fmt.Println(data)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Try programiz.pro")

	msgCh := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	emailList := []string{"abc", "def", "xyz"}

	go generateEmailWithMessage(msgCh, emailList, "message")

	go printSendMessage(msgCh, &wg)

	wg.Wait()

	fmt.Println("Task done")

}
