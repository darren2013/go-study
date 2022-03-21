package main

import (
	"fmt"
)

func main() {
	fmt.Println("a")

	//defer 将执行放入defer调用栈中，执行顺序后进先出
	defer fmt.Println("b")
	defer fmt.Println("c")
}
