package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[0:len(s1)]

	s3 := s1[:]

	//slice是一个引用类型
	s4 := s1

	fmt.Println(s2, s3, s4)
}
