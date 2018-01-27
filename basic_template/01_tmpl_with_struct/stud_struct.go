package main

import (
	"log"
	"os"
	"text/template"
)

type Student struct {
	Name string
}

func main() {
	s := Student{"Ali rahman"}

	//“New” allocates a new template with the given name.
	tmpl := template.New("test")

	// “Parse” parses a string into a template.
	tmpl, err := tmpl.Parse("Hello {{.Name}}")
	if err != nil {
		log.Fatal("Parse:", err)
		return
	}

	// used the predefined variable “os.Stdout” which refers to the standard output to print out the merged data — “os.Stdout” implements “io.Writer”.
	err1 := tmpl.Execute(os.Stdout, s)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}

}
