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

//GO语言所有传递，都是值拷贝，如果想要引用，可以使用指针
func A(p person) {
	p.Age = 88
}
