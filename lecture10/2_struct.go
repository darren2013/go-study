package main

import (
	"fmt"
)

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
