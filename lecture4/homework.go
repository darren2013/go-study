package main

import (
	"fmt"
)

const (
	//_         = iota
	B float64 = 1 << (iota * 10)
	KB
	GB
	TB
	PB
)

func main() {
	fmt.Println(KB)
}
