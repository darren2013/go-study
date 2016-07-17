package main

import (
	"fmt"
)

func main() {
	var m map[int]string

	m = make(map[int]string)

	//另一种方式简介定义
	m2 := make(map[int]string)

	//元素存
	m[1] = "ok"
	m[2] = "bad"

	//取元素
	a := m[1]

	//删除键为1的元素
	delete(m, 1)

	fmt.Println(m, m2, a)

	//元素不存在时，取得值为空
	b := m[3]
	fmt.Println("-", b+"-")
}
