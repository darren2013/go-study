package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 10)

	/*	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()*/

	//可以使用timer,设定通道读取超时
	timer := time.NewTimer(5 * time.Second)

	fmt.Println("start ", time.Now().String())

	//通道里没有数据读取时，会阻塞
	select {
	case val := <-ch:
		fmt.Println("receive ", val)
	case ct := <-timer.C:
		fmt.Println("current time", ct)

	}

	fmt.Println("end in 20 seconds")
	time.Sleep(20 * time.Second)
}
