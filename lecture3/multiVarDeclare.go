package main

import (
	"fmt"
)

func main() {
	//多个变量声明
	var a, b, c, d, e int
	//多个变量赋值
	a, b, c, d, e = 1, 2, 3, 4, 5
	fmt.Println(a, b, c, d, e)

	//多个变量声明并同时赋值
	var f, g, h int = 6, 7, 8
	fmt.Println(f, g, h)

	var i, j, k = 9, 10, 11
	fmt.Println(i, j, k)

	l, m, n := 12, 13, 123
	fmt.Println(l, m, n)
}
