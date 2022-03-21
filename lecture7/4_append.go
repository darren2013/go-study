package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1)
	fmt.Println("s2 value,cap", s2, cap(s2))
	//append以后，s2底层数组会变成一个新的数组
	s2 = append(s2, 6, 7, 8, 10)
	fmt.Println(s1, s2)
	s1[0] = 9
	fmt.Println(s1, s2)
}
