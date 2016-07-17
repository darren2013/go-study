package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println(a)
	A(&a)
	fmt.Println(a)
}

//通过指针方式，传递地址
func A(a *int) {
	*a = 2
}
