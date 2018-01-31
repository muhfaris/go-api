package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/muhfaris/goFun/basic_serialdata/01_go_mysql_rest/database"

	_ "github.com/go-sql-driver/mysql"
)

type ServeMux struct {
	// contains filtered or unexported fields
}

// DB initialize of sql.DB
var DB *sql.DB

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/users", GetUsers)
	mux.HandleFunc("/users/", GetUserId)
	http.ListenAndServe(":9393", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		Access Data To 
		Get all User : /users
		Get User by Id : /user/{id}
		Post new user : /users
		Delete the user by id : /users/{id}
		`)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dbdata := database.GetDataUsers()
		json.NewEncoder(w).Encode(dbdata)
	}
	if r.Method == "POST" {
		data := &database.RetriveUser{
			Name:  r.FormValue("Name"),
			Email: r.FormValue("Email"),
		}
		dbdata := database.SaveDataUser(data)
		fmt.Println(dbdata)
		json.NewEncoder(w).Encode(dbdata)
	}

}

func GetUserId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	if r.Method == "GET" {
		dbdata := database.GetDataUserId(id)
		json.NewEncoder(w).Encode(dbdata)
	}
	if r.Method == "DEL" {
		dbdata := database.DeleteDataUser(id)
		json.NewEncoder(w).Encode(dbdata)
	}

}
