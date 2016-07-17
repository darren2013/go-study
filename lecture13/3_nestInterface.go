package main

import (
	"fmt"
)

//接口嵌入
type USB interface {
	Name() string
	Connetor
}

type Connetor interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func main() {
	a := PhoneConnector{name: "phoneConnector"}
	//a.Name()
	//a.Connect()
	a.name = "phone"
	diconnect(a)
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("connect")
}

func diconnect(a USB) {
	if pc, ok := a.(PhoneConnector); ok {
		fmt.Println("diconnect device", pc.name)
		return
	}

	fmt.Println("unknown devicen..")
}
