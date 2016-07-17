package main

import (
	"fmt"
)

type empty interface{}

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
	pc := PhoneConnector{name: "phoneConnector"}

	var a Connetor
	//发生了拷贝
	a = Connetor(pc)
	a.Connect()

	pc.name = "phone"
	a.Connect()

}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println(pc.name, "connect")
}

func diconnect(a interface{}) {

	switch v := a.(type) {
	case PhoneConnector:
		fmt.Println("diconnect device", v.name)
	default:
		fmt.Println("unknown devicen..")
	}

}
