/*
Template_PrintPerson

本程序使用template渲染html来输出person结构
*/
package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

const template_code = `
	<div>
		<div>The name is {{.Name}}</div>
		<div>The age is {{.Age}}</div>
		<div>The emails are:</div>
		{{range .Emails}}
		<div>The emails are: {{.}}</div>
		{{end}}
		<div>The jobs are:</div>
		{{with .Jobs}}
			{{range .}}
				<div>The employer is: {{.Employer}}, the role is {{.Role}}</div>
			{{end}}
		{{end}}
	</div>
`

func main() {
	job1 := Job{Employer: "Monash", Role: "Honorary"}
	job2 := Job{Employer: "Box Hill", Role: "Head of HE"}

	person := Person{
		Name:   "Jan",
		Age:    50,
		Emails: []string{"xxxx@gmail.com", "yyyy@gmail.com"},
		Jobs:   []*Job{&job1, &job2},
	}

	t := template.New("Person template")
	t, err := t.Parse(template_code)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
