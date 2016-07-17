package main

import (
	"fmt"
)

func main() {
	s := []string{"ok", "bad", "good", "well"}

	//i是slice的索引，v是value,注意这里的v是值拷贝，如果想要修改，采取s[i]="hello"
	for i, v := range s {
		fmt.Println(i, v)
		s[i] = v + "-update"
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
