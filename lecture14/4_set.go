package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	u := User{23, "jack", 22}
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	fmt.Println(v.Kind())

	if v.Kind() != reflect.Ptr && !v.CanSet() {
		fmt.Println("不支持赋值")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")

	if !f.IsValid() {
		fmt.Println("字段名字不存在")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("Ben")
	}

}
