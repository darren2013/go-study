package main

import (
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func main() {
	p := person{Name: "张三", Age: 34}
	p.Contact.City = "hz"
	p.Contact.Phone = "777888"
	fmt.Println(p)
}
