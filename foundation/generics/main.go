package main

import "fmt"

func main() {
	clients := map[string]int{"José": 1000, "João": 5000}
	secondClients := map[string]float64{"Maria": 999.99, "Rafael": 5555.555}
	otherClients := map[string]MyNumber{"Maria": 777, "Rafael": 3333}

	fmt.Println(Sum(clients))
	fmt.Println(Sum(secondClients))
	fmt.Println(Subtraction(otherClients, 1000))
	fmt.Println(Compare(1000.0, 1000.2))
}

type MyNumber int

type CustomNumber interface {
	~int | ~float64
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func Sum[T int | float64](m map[string]T) T {
	var result T

	for _, v := range m {
		result += v
	}

	return result
}

func Subtraction[T CustomNumber](m map[string]T, subtrahend T) T {
	var result T

	for _, v := range m {
		result += v
	}

	return result - subtrahend
}
