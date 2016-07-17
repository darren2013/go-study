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
	u := User{2, "Lisi", 34}
	//Info(&u)
	Info(u)
}

func (user User) hello() {
	fmt.Println("hello", user.Name)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("type is not struct,skip it")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("fields")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fieldVal := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, fieldVal)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}
