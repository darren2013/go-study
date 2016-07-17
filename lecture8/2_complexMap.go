package main

import (
	"fmt"
)

func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)

	m[1] = make(map[int]string)
	m[1][2] = "ok"

	//函数有多返回值
	b, exists := m[2][1]

	if !exists {
		m[2] = make(map[int]string)
	}

	//m[2]对应的map未不初始化会报错
	m[2][1] = "bad"

	b, exists = m[2][1]
	fmt.Println(m, b, exists)
}
