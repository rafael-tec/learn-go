package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "Word"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
