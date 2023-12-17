package main

import "fmt"

type ID int
type name string

var (
	i ID   = 999
	n name = "Test"
)

func main() {
	fmt.Printf("type is %T\nvalue is %v\n", i, i)
	fmt.Printf("type is %T\nvalue is %v\n", n, n)
}
