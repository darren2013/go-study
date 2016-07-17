package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
}

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
