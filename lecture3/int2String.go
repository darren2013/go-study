package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := string(a)
	fmt.Println(b)

	c := strconv.Itoa(a)
	fmt.Println(c)

	//字符串转换为整形
	var d string = "10"
	e, error := strconv.Atoi(d)

	fmt.Println(error)
	fmt.Println(e)
}
