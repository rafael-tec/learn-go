package main

import "fmt"

func main() {
	m := map[string]any{"Jose": true, "Maria": 200.99, "Joao": "anything"}

	fmt.Println(m["Jose"])
	fmt.Println(m["Maria"])

	delete(m, "Jose")

	fmt.Println(m)
}
