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

type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User: User{1, "jone", 34}}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(0))

	//取user id
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
}
