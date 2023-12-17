package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 1; i <= len(numbers); i++ {
		fmt.Printf("%d, ", i)
	}

	fmt.Println("\n-------------------------------")

	for _, v := range numbers {
		fmt.Printf("%d, ", v)
	}

	fmt.Println("\n-------------------------------")

	i := 0
	for i < 10 {
		fmt.Printf("%d, ", i+1)
		i++
	}

	fmt.Println("\n-------------------------------")

	for {
		fmt.Println("In looping...")
	}
}