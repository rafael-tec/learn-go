package main

import "fmt"

var (
	b bool
	i int
	s string
	f float64
)

func main() {
	test := "a" // infered type
	test = "d"

	fmt.Println(test)
	fmt.Printf("boolean: %t\nint: %d\nstring: %s\nfloat: %f\n", b, i, s, f)
}
