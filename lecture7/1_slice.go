package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	//:后面不包含
	s1 := a[5:9]
	fmt.Println(s1)

	//取第5个元素以后的元素
	s2 := a[5:len(a)]
	fmt.Println(s2)
	s3 := a[5:]
	fmt.Println(s3)

	//取前5个元素
	s4 := a[:5]
	fmt.Println(s4)

	b := [3]int{2, 3, 12}
	s5 := b[:3]
	//c := [...]int{13, 14}
	s5 = append(s5, 13)
	fmt.Println(s5)

}
