package main

import "fmt"

// thread 1
func main() {
	c := make(chan string)

	// thread 2
	go func() {
		c <- "Hello word!"
	}()

	// thread 1
	msg := <-c
	fmt.Println(msg)
}
