package main

import (
	"fmt"
)

func main() {
	//当接口存储的类型和对象都为nil时，才是nil
	var a interface{}
	fmt.Println(a == nil)

	var p *int = nil
	a = p
	fmt.Println(a == nil)
}
