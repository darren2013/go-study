package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//waitBySleep()
	//waitByChannel()

	waitByGP()
}

func waitBySleep() {

	for j := 0; j < 1000; j++ {
		go fmt.Println(j)
	}

	time.Sleep(time.Second)
}

func waitByChannel() {
	size := 100
	barrier := make(chan bool, size)

	for i := 0; i < size; i++ {

		go func(j int) {
			fmt.Println(j)
			barrier <- true
		}(i)
	}

	for i := 0; i < size; i++ {
		<-barrier
		fmt.Println("finish ", i)
	}
}

func waitByGP() {
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
