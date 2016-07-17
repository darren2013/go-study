package main

import (
	"fmt"
)

func main() {
	//最后一个元素是1，其他元素都为0，长度是20的数组
	a := [20]int{19: 1}

	fmt.Println(a)

	//数组长度可以用...代替
	b := [...]int{2, 3, 4, 5, 6}
	fmt.Println(b)

	//会创建一个长度为100数组
	c := [...]int{99: 1}
	fmt.Println(c)
}
