package main

import (
	"fmt"
)

func main() {
	//slice里存放map
	sm := make([]map[int]string, 6)

	//注意key v都是拷贝，下面操作不会改变slice里的额值
	for _, v := range sm {
		v = make(map[int]string)
		v[1] = "ok"
		fmt.Println(v)
	}

	fmt.Println(sm)

	for i := range sm {
		sm[i] = make(map[int]string)
		sm[i][1] = "ok"
		//fmt.Println(v)
	}

	fmt.Println(sm)
}
