package main

import (
	"fmt"
)

func main() {
	fs := [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)

		//闭包里引用的，是变量的地址
		defer func() {
			fmt.Println("defer closure i=", i)
		}()

		fs[i] = func() {
			fmt.Println("closure i=", i)
		}
	}

	for _, f := range fs {
		f()
	}
}
