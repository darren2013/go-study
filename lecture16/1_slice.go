package main

import (
	"fmt"
)

func fill(s []int) {
	fmt.Printf("fill before %p %v\n", s, s)
	s = append(s, 3, 4, 5)
	fmt.Printf("fill %p %v\n", s, s)

}
func main() {
	s := make([]int, 0, 20)
	fmt.Printf("%p\n", s)
	fill(s)
	fmt.Printf("main %p %v\n", s, s)
}
