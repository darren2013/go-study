package main

import (
	"fmt"
	"time"
)

func main() {
	//设置缓存以后channel是异步的
	c := make(chan bool, 1)

	go func() {
		fmt.Println("go go....")
		<-c

		fmt.Println("read from channel over...")
	}()

	//c <- true

	time.Sleep(2 * time.Second)
}
