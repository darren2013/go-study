package main

import (
	"fmt"
)

const (
	a int = 1
	b
	//iota组中每定义一个常量，自动递增1
	c = iota
	d
	e = iota
	f
)

//遇到const关键字后，iota被重置为0
const (
	g = iota
)

func main() {
	fmt.Println(a, b, c, d, e, f, g)
}
