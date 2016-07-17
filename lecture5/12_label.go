package main

import (
	"fmt"
)

func main() {

LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			//改变执行位置，会无限循环
			goto LABEL
		}
	}
}
