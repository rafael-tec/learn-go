package main

import "fmt"

func main() {
	number := 10

	fmt.Println("Value: ", number)
	fmt.Println("Memory Address: ", &number)

	var otherNumber *int = &number

	fmt.Println("\nPointer")
	fmt.Println("Value: ", otherNumber)
	fmt.Println("Value dereference: ", *otherNumber)
	fmt.Println("Memory Address: ", &otherNumber)
}
