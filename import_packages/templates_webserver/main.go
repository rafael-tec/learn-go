package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	courses := Courses{
		Course{"Golang", 31},
		Course{"Kotlin", 40},
		Course{"Python", 30},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.New("content.html").ParseFiles(templates...))

		err := template.Execute(w, courses)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
