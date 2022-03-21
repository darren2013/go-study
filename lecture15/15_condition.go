package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := Queue{
		q:   []string{},
		con: sync.NewCond(&sync.Mutex{}),
	}

	go func() {

		for {
			queue.Enqueue("a" + time.Now().String())
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		queue.Dequeue()
		time.Sleep(time.Second)
	}

}

type Queue struct {
	q   []string
	con *sync.Cond
}

func (queue *Queue) Enqueue(item string) {
	queue.con.L.Lock()
	defer queue.con.L.Unlock()

	queue.q = append(queue.q, item)
	fmt.Printf("putting %s item to queue ,notifyall \n", item)
	queue.con.Broadcast()
}

func (queue *Queue) Dequeue() string {

	queue.con.L.Lock()
	defer queue.con.L.Unlock()

	for len(queue.q) == 0 {
		fmt.Println("no data avaiable,wait")
		queue.con.Wait()
	}
	fmt.Println("queue current size ", len(queue.q))
	result := queue.q[0]
	queue.q = queue.q[1:]

	return result
}
