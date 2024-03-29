package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)

	for i := 0; i < 10; i++ {
		go GO(c, i)
	}

	for i := 0; i < 10; i++ {
		//<-c
		fmt.Println(i, <-c)
	}
}

func GO(c chan bool, index int) {
	a := 1

	for i := 0; i < 100000; i++ {
		a += i
	}

	fmt.Println(index, a)
	c <- true
}
