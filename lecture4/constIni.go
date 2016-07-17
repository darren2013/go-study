package main

import (
	"fmt"
)

//常量组里，如果不提供初始化，则使用上行表达式的值
const (
	a = 1
	b
	c
)

func main() {
	fmt.Println(a, b, c)
}
