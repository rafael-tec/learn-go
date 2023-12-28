package main

import "fmt"

func main() {
	ch := make(chan int)

	go publish(ch)
	consume(ch)
}

func consume(ch chan int) {
	for v := range ch {
		fmt.Printf("Received: %d\n", v)
	}
}

func publish(ch chan int) {
	for i := 1; i < 10; i++ {
		fmt.Printf("Sended: %d\n", i)
		ch <- i
	}
	close(ch)
}
