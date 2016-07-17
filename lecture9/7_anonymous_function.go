package main

import (
	"fmt"
)

func main() {
	//匿名函数
	a := func() {
		fmt.Println("function A")
	}

	a()
}
