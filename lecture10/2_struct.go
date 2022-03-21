package main

import (
	"fmt"
)

//结构体的属性，定义时最好大小，在外面能够直接访问
type person struct {
	Name string
	Age  int
}

func main() {

	p := person{
		Name: "张三",
		Age:  34,
	}

	//另外一种初始化类型
	b := person{}
	b.Age = 34
	b.Name = "李四"

	fmt.Println(p, b)
}
