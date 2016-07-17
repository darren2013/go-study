package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

func B() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in b")
		}
	}()
	panic("panic in B")
}

func C() {
	fmt.Println("Func C")
}
