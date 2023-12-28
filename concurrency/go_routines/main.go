package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)

		if i == 2 {
			time.Sleep(100 * time.Second)
		}
	}
	fmt.Println()
}

func main() {
	go task("A1")
	go task("B2")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
		fmt.Println()
	}()

	time.Sleep(11 * time.Second)
}
