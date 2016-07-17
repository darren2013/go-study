package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	//初始化就取地址，结构体这样就支持引用，建议这种方式
	p := &person{}
	p.Name = "李四"
	p.Age = 34
	fmt.Println(p)
	//结构体进行的是值拷贝，而非引用
	A(p)
	B(p)
	fmt.Println(p)
}

func A(p *person) {
	p.Age = 88
}

func B(p *person) {
	p.Age = 199
}
