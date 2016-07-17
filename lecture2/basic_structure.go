//package声明一定要放在非注释的首行代码最前面
package main

//为包建立别名
import std "fmt"

//常量组
const (
	PI     = 3.14
	const1 = 2.34
	const2 = 4
)

//全局变量的声明与赋值
var (
	name  = "gopher"
	name1 = "test1"
	name2 = "test2"
	name3 = 3
)

//一般类型声明
type (
	a int
	b float32
	c string
	d byte
)

type gopher struct{}

type golang interface{}

func main() {
	std.Println("ddhello,world,你好世界啊啊啊啊！！")
}
