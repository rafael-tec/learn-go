package main

import "fmt"

func main() {
	var value interface{} = "Jose Silva"
	fmt.Printf("Value is: %s\n", value.(string))

	res, ok := value.(int)
	fmt.Printf("Value is %d and parse is: %v", res, ok)

	res2 := value.(int)
	fmt.Println(res2)
}
