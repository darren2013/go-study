package main

import (
	"fmt"
)

func main() {
	a := 5

	switch {
	case a > 0:
		fmt.Println("a > 0")
		fallthrough
	case a > 1:
		fmt.Println("a > 1")
	default:
		fmt.Println("NONE")
	}
}
