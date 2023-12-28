package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(9)

	fmt.Println("Starting...")

	go publish(ch)
	go consume(ch, &wg)

	fmt.Println("Waiting other threads...")

	wg.Wait()
}

func consume(ch chan int, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Printf("Received: %d\n", v)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 1; i < 10; i++ {
		fmt.Printf("Sended: %d\n", i)
		ch <- i
	}
	close(ch)
}
