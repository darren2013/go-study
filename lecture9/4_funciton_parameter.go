package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2}
	fmt.Println(s)
	A(s)
	fmt.Println(s)
}

//拷贝的是地址
func A(a []int) {
	a[0] = 3
	a[1] = 4

}
