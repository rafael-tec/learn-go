package main

import (
	"errors"
	"fmt"
)

func main() {
	division, _ := division(1, 2)
	fmt.Printf("result division: %.2f\n", division)

	sum := sum(1, 2, 3, 4, 5)
	fmt.Printf("result sum: %d\n", sum)

	nextInt := intSeq()
	fmt.Printf("result nextInt: %d\n", nextInt())
	fmt.Printf("result nextInt: %d\n", nextInt())
	fmt.Printf("result nextInt: %d\n", nextInt())
}

func division(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisor cannot be zero")
	}

	return float64(a) / float64(b), nil
}

func sum(numbers ...int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
