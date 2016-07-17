package main

import (
	"fmt"
)

func main() {
	a := 1

	//无限循环
	for {
		if a > 5 {
			break
		}
		a++
	}

	fmt.Println(a)
}
