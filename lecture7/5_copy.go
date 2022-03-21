package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}

	//copy不是追加，而是对应位置覆盖
	copy(s2, s1)
	fmt.Println(s2)
	//copy(s1[1:3], s2[:2])
	fmt.Println(s1)

}
