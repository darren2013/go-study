package main

import (
	"fmt"
)

func main() {
	a := 1
	//++只能作为一个独立语句存在
	//var b = a++
	a++
	//++只能在右侧
	//++a
	var b = a

	fmt.Println(b)
}
