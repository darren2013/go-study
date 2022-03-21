package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go func() {
		for {
			select {
			case v := <-done:
				fmt.Println("stop go routine...", v)
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)

	close(done)

	time.Sleep(5 * time.Second)
}
