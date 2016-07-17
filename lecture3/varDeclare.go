package main

import (
	"fmt"
)

func main() {
	//变量声明
	var a int
	//变量赋值
	a = 1
	fmt.Println(a)

	//变量声明并同时赋值
	var b int = 43
	fmt.Println(b)

	//省略变量类型，由系统推断
	var c = 556
	fmt.Println(c)

	//变量声明与赋值简写
	d := 100
	fmt.Println(d)

}
