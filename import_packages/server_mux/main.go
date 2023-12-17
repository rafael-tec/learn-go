package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHelloWord)
	mux.Handle("/blog", Blog{Title: "My blog"})

	http.ListenAndServe(":8080", mux)
}

func GetHelloWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello word!"))
}

type Blog struct {
	Title string
}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.Title))
}
