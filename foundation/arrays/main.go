package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3}
	arr[3] = 4
	arr[4] = 5

	fmt.Println(arr)
}
