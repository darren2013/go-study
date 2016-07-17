package main

import (
	"fmt"
)

type TZ int

func main() {
	var a TZ
	a.increase(100)
	fmt.Println(a)
}

func (tz *TZ) increase(num int) {
	//注意这里TZ必须做强制类型转换
	*tz += TZ(num)
}
