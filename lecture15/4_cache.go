package main

import (
	"fmt"
)

func main() {
	//设置缓存以后channel是异步的
	c := make(chan bool, 1)

	go func() {
		fmt.Println("go go....")
		<-c
	}()

	c <- true
}
