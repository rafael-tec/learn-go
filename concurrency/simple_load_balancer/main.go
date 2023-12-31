package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	qtWorker := 100

	for i := 1; i <= qtWorker; i++ {
		go worker(i, data)
	}

	for i := 1; i <= 1000; i++ {
		data <- i
	}
}

func worker(workerId int, data <-chan int) {
	for v := range data {
		fmt.Printf("Worker %d received request %d\n", workerId, v)
		time.Sleep(time.Second)
	}
}
