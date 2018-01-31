package database

import (
	"database/sql"
	"fmt"
)

const (
	username = "dev"
	password = "dev123"
	dbname   = "db1"
	server   = "localhost"
	port     = 3306
)

// User for get user from db
type User struct {
	Id    int
	Name  string
	Email string
}

type RetriveUser struct {
	Name  string
	Email string
}

func Connect() (*sql.DB, error) {
	dbtext := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, server, port, dbname)
	db, err := sql.Open("mysql", dbtext)

	if err == nil {

	}
	return db, nil
}

func GetDataUsers() (data []*User) {
	DB, err := Connect()
	var u User
	defer DB.Close()
	if err == nil {
		// Read data
		readdb, _ := DB.Query("select * from user")
		for readdb.Next() {
			readdb.Scan(&u.Id, &u.Name, &u.Email)
			data = append(data, &User{
				Id:    u.Id,
				Name:  u.Name,
				Email: u.Email,
			})
		}
	} else {
		panic(err.Error())
	}
	return
}

func GetDataUserId(id string) (data *User) {
	DB, err := Connect()
	var u User
	defer DB.Close()
	if err == nil {
		readdb, _ := DB.Query("select * from user where id=?", id)
		for readdb.Next() {
			readdb.Scan(&u.Id, &u.Name, &u.Email)
			data = &User{
				Id:    u.Id,
				Email: u.Email,
				Name:  u.Name,
			}
		}
	} else {
		panic(err.Error())
	}
	return
}

func SaveDataUser(u *RetriveUser) (result string) {
	DB, err := Connect()
	defer DB.Close()

	if err == nil {
		_, err := DB.Query("insert into user (Name, Email) values(?,?)", u.Name, u.Email)
		if err != nil {
			result = err.Error()
		} else {
			result = "Success, Data is save!"
		}
	} else {
		panic(err.Error())
	}
	return
}

func DeleteDataUser(id string) (result string) {
	DB, err := Connect()
	defer DB.Close()

	if err == nil {
		_, err := DB.Query("Delete from user where id=?", id)
		if err != nil {
			result = err.Error()
		} else {
			result = "Success, Data is deleting!"
		}
	}
	return
}
