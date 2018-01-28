package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

// StaticRoot and StaticUrl is default value of static file
const (
	StaticUrl  string = "/static/"
	StaticRoot string = "static"
)

// Context is variabel value
type Context struct {
	Title  string
	Static string
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/contact", ContactHandler)
	http.HandleFunc(StaticUrl, StaticHandler)
	http.ListenAndServe(":9393", nil)
}

// HomeHandler is controller for page index
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home Page")
	Context := Context{Title: "welcome to home!"}
	RenderTemplate(w, "home", Context)
}

// AboutHandler is controller for page about
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("About Page")
	Context := Context{Title: "About Page!"}
	RenderTemplate(w, "about", Context)
}

// ContactHandler is controller for page contact
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Contact Page")
	Context := Context{Title: "Contact Page!"}
	RenderTemplate(w, "contact", Context)
}

// RenderTemplate function for load templates
func RenderTemplate(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = StaticUrl
	tmpllist := []string{"templates/base.html", fmt.Sprintf("templates/%s.html", tmpl)}

	for _, layuot := range tmpllist {
		fmt.Println(layuot)
	}

	t, err := template.ParseFiles(tmpllist...)
	if err != nil {
		log.Print("Template parsing error:", err)
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Print("Template executing error:", err)
	}
}

// StaticHandler function to load static file
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	staticfile := r.URL.Path[len(StaticUrl):]
	fmt.Println(staticfile)
	if len(staticfile) != 0 {
		f, err := http.Dir(StaticRoot).Open(staticfile)

		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, r, staticfile, time.Now(), content)
			return
		}
	}
	http.NotFound(w, r)
}
