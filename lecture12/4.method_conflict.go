package main

import (
	"fmt"
)

type A struct {
	B
	C
}

type B struct {
}

type C struct {
}

func main() {
	a := A{}
	a.Print()
	a.B.Print()
	a.C.Print()
}

func (a *A) Print() {
	fmt.Println("A.....")
}

func (b *B) Print() {
	fmt.Println("B.....")
}

func (c *C) Print() {
	fmt.Println("C.....")
}
