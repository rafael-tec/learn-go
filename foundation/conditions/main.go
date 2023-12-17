package main

import "fmt"

func main() {
	number := 10
	valid := true

	if number < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Equal to or greater than 10")
	}

	if number < 100 && valid {
		fmt.Println("Number is valid and less than 10")
	}

	switch number {
	case 1:
		fmt.Println("Equal to 1")
	case 2:
		fmt.Println("Equal to 2")
	default:
		fmt.Println("Numbers isn't valid")
	}
}