package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 4; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)

		if i == 2 {
			time.Sleep(30 * time.Second)
		}

		wg.Done()
	}
	fmt.Printf("Finish task %s\n", name)
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(12)

	go task("A1", &waitGroup)
	go task("B2", &waitGroup)

	go func() {
		for i := 0; i < 4; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
		fmt.Printf("Finish task anonymous\n")
	}()

	waitGroup.Wait()
}
