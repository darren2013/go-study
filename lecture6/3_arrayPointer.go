package main

import (
	"fmt"
)

func main() {
	a := [...]int{19: 1}
	//指向数组的指针
	var p *[20]int = &a
	fmt.Println(p)

	a[1] = 2
	fmt.Println(a)
	//使用new关键字获取指向数组的指针
	p2 := new([20]int)
	//指向数组的指针可以通过下标赋值
	p2[0] = 9
	fmt.Println(p2)

	x, y := 1, 2

	//指针数组
	b := [...]*int{&x, &y}
	fmt.Println(b)

	//二维数组
	c := [2][3]int{
		{1, 2, 3},
		{4, 5, 6}}
	//结束}一定要和最后一个元素同一行
	fmt.Println(c)

}
