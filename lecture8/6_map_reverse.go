package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m)
	m2 := make(map[string]int)

	for k, v := range m {
		m2[v] = k
	}

	fmt.Println(m2)
}
