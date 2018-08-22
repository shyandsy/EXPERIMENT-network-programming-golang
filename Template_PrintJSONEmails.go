/*
Template_PrintJSONEmails

本程序演示了在template中使用条件判断(if else)以及循环(for)的用法
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
{
	"Name": "{{.Name}}",
	"Emails": [
		{{range $index, $element := .Emails}}
			{{if $index}}
				, "{{$element}}"
			{{else}}
				"{{$element}}"
			{{end}}
		{{end}}
	]
}
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
