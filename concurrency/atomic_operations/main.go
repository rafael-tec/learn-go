package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var visitsCount uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&visitsCount, 1)
		w.Write([]byte(fmt.Sprintf("You are visitor number %d\n", visitsCount)))
	})

	http.ListenAndServe(":3000", nil)
}
