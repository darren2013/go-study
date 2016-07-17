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

func (u User) Hello(msg string) string {
	fmt.Println(u.Name, msg)
	return u.Name
}

func main() {
	var any interface{}
	any = User{23, "Jack", 45}
	v := reflect.ValueOf(any)
	m := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("hi")}
	res := m.Call(args)
	fmt.Println(res)
}
