package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int, 10)

	for i := 0; i < 10; i++ {
		queue <- i
		time.Sleep(time.Second)
	}

	i := 0

	for {

		if i >= 10 {
			break
		}

		fmt.Println(<-queue)
		time.Sleep(time.Second)
		i++
	}

	fmt.Println("end.....")

}
