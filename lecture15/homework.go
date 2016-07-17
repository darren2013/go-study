package main

import (
	"fmt"
)

var c chan string

func Pinpong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From pinpong hi #%d", i)
		i++
	}
}

func main() {
	c = make(chan string)
	go Pinpong()

	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From main hello #%d", i)

		fmt.Println(<-c)
	}
}
