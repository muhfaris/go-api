package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type typePost struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type typeResponse struct {
	Message string `json:"message"`
}

func home(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "index.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	//return a 404 if template doesnt exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// if doestn directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func postsTemplate(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	msg := r.FormValue("message")

	data := &typePost{
		Title:   title,
		Message: msg,
	}
	err := saveToFile(data)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	d := typeResponse{
		Message: "success",
	}
	json, err := json.Marshal(d)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func saveToFile(param *typePost) error {
	json, err := json.Marshal(param)
	if err != nil {
		http.StatusText(500)
		return err
	}
	d1 := []byte(json)
	err = ioutil.WriteFile("files/data.json", d1, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	dt := http.FileServer(http.Dir("files"))
	http.Handle("/data/", http.StripPrefix("/data/", dt))

	http.HandleFunc("/", home)
	http.HandleFunc("/posts", postsTemplate)

	fmt.Println("Listening ...")
	http.ListenAndServe(":7171", nil)
}
