package main

import (
	"fmt"
)

func main() {
	//go中不存在隐士数据类型转换
	//转换只发生在两种相互兼容的类型之间
	//go一种类型安全的语言
	var a float32 = 100.1

	b := int(a)
	fmt.Println(b)

	//下面类型会转换失败
	//var c bool = false
	//d := int(c)
	//fmt.Println(d)
}
