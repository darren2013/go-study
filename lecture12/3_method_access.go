package main

import (
	"fmt"
)

type A struct {
	name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a)
}

//方法是可以访问私有字段的
func (a *A) Print() {
	a.name = "张三"
}
