package main

import (
	"fmt"
)

//定义常量，常量是在编译时就确定了，常量定义的名字一般大写
//大写意味着访问类型是private,如果想访问类型是public,变量名字以_开头

const A int = 1
const B = 'A'

//定义private类型的常量
const _c int = 234

const (
	TEXT = "123"
	//只能使用内置函数
	LENGTH = len(TEXT)
	D      = B * 20
)

func main() {
	fmt.Println(A, B, _c, TEXT, LENGTH)
}
