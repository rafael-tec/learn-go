package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	courses := Courses{
		Course{"Golang", 31},
		Course{"Kotlin", 40},
		Course{"Python", 30},
	}

	t := template.New("template.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles("template.html"))

	err := t.Execute(os.Stdout, courses)
	if err != nil {
		panic(err)
	}
}
