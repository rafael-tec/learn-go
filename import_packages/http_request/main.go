package main

import (
	"io"
	"net/http"
	"fmt"
)

func main() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response body: %s\n", body)
}