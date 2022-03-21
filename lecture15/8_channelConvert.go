package main

func main() {

	ch := make(chan int)

	go produce(ch)

	go consum(ch)
}

//只能发送消息
func produce(c chan<- int) {
	for {
		c <- 1
	}
}

//只能接收消息
func consum(c <-chan int) {
	for {
		<-c
	}
}
