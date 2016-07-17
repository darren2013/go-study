package main

import (
	"fmt"
)

type human struct {
	Sex int
}

//不支持继承，只支持组合方式
type person struct {
	human
	Name string
	Age  int
}

func main() {
	p := person{Name: "张三", Age: 23, human: human{Sex: 1}}
	p.Sex = 0
	fmt.Println(p)
}
