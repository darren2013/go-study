package main

import (
	"fmt"
)

type A struct {
	B
	C
	Name string
}

type B struct {
	Name string
}

type C struct {
	Name string
}

func main() {
	a := A{B: B{Name: "李四"}, C: C{Name: "王五"}}
	a.Name = "A"
	//之前go版本有相同级别时，有相同的name时编译器会报错的
	fmt.Println(a.Name)
	fmt.Println(a.B.Name)
	fmt.Println(a.C.Name)
}
