package main

import (
	"sync"
	"time"
)

func main() {
	/*conflictMap := map[int]int{}
	conflictMap[1] = 999
	unsafeMap(conflictMap)
	fmt.Println(conflictMap)*/

	safeWrit()

	time.Sleep(10 * time.Second)

}

func safeWrit() {
	s := SafeMap{
		safeMap: map[int]int{},
		Mutex:   sync.Mutex{},
	}

	for i := 0; i < 100; i++ {
		go func() {
			s.writer(1, 2)
		}()
	}
}

func unsafeMap(conflictMap map[int]int) {

	for i := 0; i < 10; i++ {
		go func() {
			conflictMap[1] = i
		}()
	}
}

type SafeMap struct {
	safeMap map[int]int
	sync.Mutex
}

func (safeMap *SafeMap) read(key int) (int, bool) {
	safeMap.Mutex.Lock()
	defer safeMap.Mutex.Unlock()

	result, exist := safeMap.safeMap[key]

	return result, exist
}

func (safeMap *SafeMap) writer(k, v int) {
	safeMap.Mutex.Lock()
	defer safeMap.Mutex.Unlock()

	safeMap.safeMap[k] = v
}
