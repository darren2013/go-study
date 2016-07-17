package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p \n", s1)

	//追加后，小于容量，则使用原始数组
	s1 = append(s1, 12, 1, 3)
	fmt.Printf("%v,%p\n", s1, s1)

	//大于容量，重新分配一个数组
	s1 = append(s1, 12, 1, 3)
	fmt.Printf("%v,%p", s1, s1)
}
