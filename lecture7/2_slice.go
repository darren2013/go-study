package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3, 20)
	fmt.Println(s, len(s), cap(s))

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	s2 := a[2:5]
	//cap会一直持续到底层数组的底部
	fmt.Println(len(s2), cap(s2))

	fmt.Println(string(s2))

	//reslice,索引要重新计算
	s3 := s2[1:3]
	fmt.Println(string(s3))
}
