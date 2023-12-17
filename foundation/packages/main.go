package main

import (
	"fmt"
	"packages/math"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(math.Sum(10, 20))
	fmt.Println(uuid.New())
}
