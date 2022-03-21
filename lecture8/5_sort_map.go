package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{9f: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}
	s := make([]int, len(m))
	i := 0

	for k, _ := range m {
		s[i] = k
		i++
	}

	sort.Ints(s)
	fmt.Println(s)
}
