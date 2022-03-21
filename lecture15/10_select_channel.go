package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	messages := make(chan string, 10)
	done := make(chan bool)

	defer close(messages)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		time.NewTimer()
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("receive message ", <-messages)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		messages <- string("index" + strconv.Itoa(i))
	}

	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit....")
}
