package main

import (
	"fmt"
	"sync"
)

type SliceNum []int

func createSliceNum() SliceNum {
	return make(SliceNum, 0)
}

func (sliceNum *SliceNum) Add(item int) *SliceNum {

	*sliceNum = append(*sliceNum, item)
	fmt.Println("add num ", item)
	return sliceNum
}

func main() {
	var once sync.Once
	sliceNum := createSliceNum()

	once.Do(func() {
		sliceNum.Add(123)
	})

	once.Do(func() {
		sliceNum.Add(12)
	})

	once.Do(func() {
		sliceNum.Add(13)
	})

}
