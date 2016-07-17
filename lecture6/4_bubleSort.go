package main

import (
	"fmt"
)

func main() {
	a := [...]int{92, 15, 100, 24, 56, 32}

	fmt.Println(a)
	length := len(a)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}

	fmt.Println(a)
}
