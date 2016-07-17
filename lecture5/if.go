package main

import (
	"fmt"
)

func main() {

	/*
		1.if表达式后面没有()；
		2.{必须和if在同一行
		3.a是局部变量，只能if语句块和else语句块中访问
	*/
	if a := 1; a > 1 {
		fmt.Println("a > 1")
	} else {
		fmt.Println("a = ", a)
	}

	fmt.Println("over")
}
