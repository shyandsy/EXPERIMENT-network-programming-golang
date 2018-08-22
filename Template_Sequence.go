/*
Template_Sequence
*/
package main

import (
	"errors"
	"fmt"
	"html/template"
	"os"
)

var template_code = `
{{$comma := sequence "-" " *"}}
{{range $}}{{$comma.Next}}{{.}}{{end}}
-----------
{{$comma := sequence "" ", "}}
{{$colour := cycle "black" "white" "red"}}
{{range $}}{{$comma.Next}}{{.}} IN {{$colour.Next}}{{end}}
-----------
`

var fmap = template.FuncMap{
	"sequence": sequenceFunc,
	"cycle":    cycleFunc,
}

func main() {
	t, err := template.New("").Funcs(fmap).Parse(template_code)
	if err != nil {
		fmt.Printf("parse error %v\n", err)
	}
	err = t.Execute(os.Stdout, []string{"a", "b", "c", "d", "e", "f"})
	if err != nil {
		fmt.Printf("execute error: %v\n", err)
	}
}

type generator struct {
	ss []string
	i  int
	f  func(s []string, i int) string
}

func (seq *generator) Next() string {
	s := seq.f(seq.ss, seq.i)
	seq.i++
	return s
}

func sequenceGen(ss []string, i int) string {
	if i >= len(ss) {
		return ss[len(ss)-1]
	}
	return ss[i]
}

func cycleGen(ss []string, i int) string {
	return ss[i%len(ss)]
}

func sequenceFunc(ss ...string) (*generator, error) {
	if len(ss) == 0 {
		return nil, errors.New("sequence must have at least one element")
	}
	return &generator{ss, 0, sequenceGen}, nil
}

func cycleFunc(ss ...string) (*generator, error) {
	if len(ss) == 0 {
		return nil, errors.New("cycle must have at least one element")
	}
	return &generator{ss, 0, cycleGen}, nil
}
