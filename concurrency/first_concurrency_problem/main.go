package main

import (
	"fmt"
	"net/http"
	"time"
)

var visitsCount = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visitsCount++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("You are visitor number %d\n", visitsCount)))
	})

	http.ListenAndServe(":3000", nil)
}
