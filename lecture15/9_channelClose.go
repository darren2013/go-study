package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	defer close(ch)

	ch <- 1

	if v, notClosed := <-ch; notClosed {
		fmt.Println(v)
	}

}
