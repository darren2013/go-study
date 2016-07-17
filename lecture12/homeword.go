package main

import (
	"fmt"
)

type A struct {
	counter int
}

func main() {
	a := A{}
	a.count()
	a.count()
	fmt.Println(a)
}

func (a *A) count() {
	a.counter += 100
}
