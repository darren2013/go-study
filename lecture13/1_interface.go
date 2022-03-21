package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
}

//结构体实现了接口的所有方法，就认为实现了该接口
type PhoneConnector struct {
	name string
}

func main() {
	var a USB
	a = PhoneConnector{name: "phoneConnector"}
	//a.Name()
	a.Connect()
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("connect")
}
