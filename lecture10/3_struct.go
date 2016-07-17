package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p := person{}
	p.Name = "李四"
	p.Age = 34
	fmt.Println(p)
	//结构体进行的是值拷贝，而非引用
	A(p)
	fmt.Println(p)
}

func A(p person) {
	p.Age = 88
}
