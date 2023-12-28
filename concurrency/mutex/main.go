package main

import (
	"fmt"
	"net/http"
	"sync"
)

var visitsCount = 0

func main() {
	m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		visitsCount++
		m.Unlock()

		w.Write([]byte(fmt.Sprintf("You are visitor number %d\n", visitsCount)))
	})

	http.ListenAndServe(":3000", nil)
}
