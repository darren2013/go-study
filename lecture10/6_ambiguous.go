package main

import (
	"fmt"
)

func main() {
	//结构体匿名创建
	a := struct {
		Name string
		Age  int
	}{
		Name: "张三",
		Age:  19,
	}
	fmt.Println(a)
}
