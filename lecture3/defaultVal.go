package main

import (
	"fmt"
	"math"
)

type (
	byte int8
	rune int32
	文本   string
)

func main() {
	var a int32
	fmt.Println("int32", a)

	var b float32
	fmt.Println("float32", b)

	var c bool
	fmt.Println("bool", c)

	var d string
	fmt.Println("string", d)

	var e [2]int
	fmt.Println(e)

	fmt.Println("maxInt32", math.MaxInt32)

	var f 文本 = "中文"
	fmt.Println("文本类型", f)

}
