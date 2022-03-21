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

type TVConnector struct {
	name string
}

func main() {
	a := PhoneConnector{name: "phoneConnector"}
	//a.Name()
	//a.Connect()
	a.name = "phone"
	var b Connetor
	b = Connetor(a)
	b.Connect()
	diconnect(a)

	tv := TVConnector{name: "TV"}
	var c USB
	c = USB(tv)
	c.Connect()
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("phone connect")
}

func (tv TVConnector) Connect() {
	fmt.Println("tv connect")
}

func (tv TVConnector) Name() string {
	fmt.Println("tv name", tv.name)
	return tv.name
}

func diconnect(a interface{}) {

	switch v := a.(type) {
	case PhoneConnector:
		fmt.Println("diconnect device", v.name)
	default:
		fmt.Println("unknown devicen..")
	}

}
