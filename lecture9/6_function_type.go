package main

import (
	"fmt"
)

func main() {
	a := A
	a()
}

//函数也作为一种类型
func A() {
	fmt.Println("function A..")
}
