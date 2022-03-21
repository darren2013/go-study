package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	A(a, b)
	fmt.Println(a, b)
}

//不定长参数列表，必须作为最后一个参数
func A(a ...int) {
	fmt.Println(a[0], a[1])
	a[0] = 3
	a[1] = 4
	fmt.Println(a)
}
