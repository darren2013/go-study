package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.print()
	fmt.Println(a)

	b := B{}
	b.print()
	fmt.Println(b)
}

//这两个方法的接受者不同，不属于方法重载，属于两个不同的方法
func (a *A) print() {
	a.Name = "张三"
	fmt.Println("A..")
}

func (b B) print() {
	b.Name = "李四"
	fmt.Println("B....")
}
