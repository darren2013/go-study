package main

import (
	"fmt"
)

func main() {

	a := 10

	//if变量声明时，遇到和外部变量同名的，外部变量会临时隐藏起来，if语句块访问的都是局部的

	if a := 1; a > 1 {
		a++
		fmt.Println(a)
	}

	fmt.Println(a)
}
