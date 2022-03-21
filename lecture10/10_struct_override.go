package main

import (
	"fmt"
)

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{Name: "张三", B: B{Name: "李四"}}
	//a.Name只能访问到A里定义的name
	fmt.Println(a.Name)
	fmt.Println(a.B.Name)
	fmt.Println(a)
}
