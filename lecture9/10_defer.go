package main

import (
	"fmt"
)

func main() {
	//defer发送严重错误时也会执行
	for i := 0; i < 3; i++ {
		defer func() {
			//闭包，引用的是变量的地址
			fmt.Println(i)
		}()
	}
}
