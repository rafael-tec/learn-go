package main

import (
	"fmt"
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	mustRenderTemplate()
	standardRenderTemplate()
	renderHtmlTemplate()
}

func renderHtmlTemplate() {
	fmt.Println("Executing render html template...")

	template := template.Must(template.New("template.html").ParseFiles("template.html"))
	courses := Courses{
		Course{"Golang", 31},
		Course{"Kotlin", 40},
		Course{"Python", 30},
	}

	err := template.Execute(os.Stdout, courses)
	if err != nil {
		panic(err)
	}
}

func mustRenderTemplate() {
	fmt.Println("Executing must render template...")

	course := Course{"Golang", 20}
	template := template.Must(
		template.New("CourseTemplate").Parse("Course: {{.Name}} - Workload: {{.Workload}}\n"),
	)

	err := template.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}

func standardRenderTemplate() {
	fmt.Println("Executing standard render template...")

	course := Course{"Golang", 20}

	template := template.New("CourseTemplate")
	template, _ = template.Parse("Course: {{.Name}} - Workload: {{.Workload}}\n")
	err := template.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
