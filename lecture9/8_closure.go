package main

import (
	"fmt"
)

func main() {
	a := closure(10)

	//下面两次调用，x的地址是一样的
	fmt.Println(a(25))
	fmt.Println(a(35))
}

//闭包
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
