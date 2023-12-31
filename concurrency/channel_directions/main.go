package main

import "fmt"

func main() {
	ch := make(chan string)

	go receive("Hello", ch)
	read(ch)
}

func receive(name string, data chan<- string) {
	// channel receive data
	data <- name
}

func read(data <-chan string) {
	// read data from channel
	value := <-data
	fmt.Println(value)
}
