package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello word!"

	showType(x)
	showType(y)
}

func showType(i interface{}) {
	fmt.Printf("Type is: %T and value is: %v\n", i, i)
}
