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

func main() {
	mustRenderTemplate()
	standardRenderTemplate()
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
