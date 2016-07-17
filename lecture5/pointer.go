package main

import (
	"fmt"
)

func main() {
	a := 1
	var p *int = &a
	//打印地址
	fmt.Println(p)
	fmt.Println("内容:", *p)
}
