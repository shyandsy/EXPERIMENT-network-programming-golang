/*
Template_PrintEmails
本程序演示如何给template增加自定义函数。
本程序为template添加了函数EmailExpander用于把所有email中的@替换为at来显示
*/
package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Person struct {
	Name   string
	Emails []string
}

const template_code = `
<div>The name is {{.Name}}</div>
<div>The email are</div>
{{range .Emails}}
<div>"{{. | emailExpand}}"</div>
{{end}}
`

func EmailExpander(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	// 查找@
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	// 替换@
	return substrs[0] + " at " + substrs[1]
}

func main() {
	person := Person{
		Name:   "Jan",
		Emails: []string{"xxx@gmail.name", "yyy@gmail.com"},
	}

	t := template.New("Person template")

	// 添加函数
	t = t.Funcs(template.FuncMap{"emailExpand": EmailExpander})

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
