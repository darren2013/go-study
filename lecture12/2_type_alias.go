package main

import (
	"fmt"
)

type TZ int

func main() {
	var tz TZ
	//method value
	tz.Print()

	//method expression
	(*TZ).Print(&tz)
}

func (tz *TZ) Print() {
	fmt.Println("tz", tz)
}
