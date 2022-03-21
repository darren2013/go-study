package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type person1 struct {
	Name string
	Age  int
}

func main() {
	p1 := person{Name: "张三", Age: 12}
	//p2 := person1{Name: "李四", Age: 34}
	//fmt.Println(p1 == p2)
	p3 := person{Name: "张三", Age: 12}
	fmt.Println(p1 == p3)

}
