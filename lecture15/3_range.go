package main

import (
	"fmt"
)

func main() {

	c := make(chan bool)

	go func() {
		fmt.Println("go go....")
		c <- true
		close(c)
	}()

	//channel关闭后，循环结束
	for v := range c {
		fmt.Println(v)
	}
}
