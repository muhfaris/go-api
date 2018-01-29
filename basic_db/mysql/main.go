package main

import (
	"database/sql"
	"log"

	"github.com/muhfaris/goFun/basic_db/mysql/database"

	_ "github.com/go-sql-driver/mysql"
)

// DB initialize of sql.DB
var DB *sql.DB

// User for get user from db
type User struct {
	Name  string
	Email string
}

func main() {
	DB, err := database.Connect()
	var u User
	var data []*User
	defer DB.Close()
	if err == nil {
		// Read data
		readdb, _ := DB.Query("select name, email from user")
		for readdb.Next() {
			readdb.Scan(&u.Name, &u.Email)
			data = append(data, &User{
				Name:  u.Name,
				Email: u.Email,
			})
		}
	} else {
		panic(err.Error())
	}

	for x, y := range data {
		log.Printf("Data ke-%d = nama (%s) / email(%s)", x, y.Name, y.Email)
	}
}
