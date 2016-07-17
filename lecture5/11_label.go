package main

import (
	"fmt"
)

func main() {
LABEL:
	for {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break LABEL
			} else {
				fmt.Println(i)
			}
		}
	}

	fmt.Println("over")
}
