/*
Template_PrintNameEmails

本程序演示如何在模板代码中使用变量
*/
package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name   string
	Emails []string
}

const template_code = `
{{$name := .Name}}
{{range .Emails}}
<div>Name is {{$name}}, Email is {{.}}</div>
{{end}}
`

func main() {
	person := Person{
		Name:   "Jan",
		Emails: []string{"xxx@gmail.name", "yyy@gmail.com"},
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
