package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const tmpl = `
The name is {{.Name}}
{{range .Emails}}
	His email id is {{.}}
{{end}}
`

func main() {
	person := Person{
		Name:   "Ali Rahman",
		Emails: []string{"test1@email.com", "tetst2@email.com", "test3@email.com"},
	}
	t := template.New("Person Template")
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	err = t.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal(err)
		return
	}
}
