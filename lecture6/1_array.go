package main

import (
	"fmt"
)

func main() {
	//定义了长度为2的int型数组，长度也作为类型一部分
	//数组为值类型
	var a [2]int
	var b [2]int

	fmt.Println(a == b)

	fmt.Println(a)
}
