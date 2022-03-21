package main

import (
	"fmt"
)

func main() {

	c := make(chan bool)

	go func() {
		fmt.Println("go go....")
		c <- true
	}()

	boolVal := <-c

	fmt.Println("end....", boolVal)
}
