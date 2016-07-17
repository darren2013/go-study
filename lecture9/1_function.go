package main

import (
	"fmt"
)

func main() {
	fmt.Println(Add(5, 4))

	c, d, e := test()
	fmt.Println(c, d, e)

	e, f, g := test2()
	fmt.Println(e, f, g)
}

func Add(x int, y int) int {
	return x + y
}

//返回值命名，return中就可以不写了
func test() (a int, b int, c int) {
	a, b, c = 2, 3, 4
	return
}

//返回值没有命名时，return返回需要指定
func test2() (int, int, int) {
	a, b, c := 1, 2, 3
	return a, b, c
}
