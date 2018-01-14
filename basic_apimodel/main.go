package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Person struct {
	ID        uint   `json : "id"`
	Firstname string `json : "firstname"`
	Lastname  string `json : "lastname"`
}

var db *gorm.DB
var err error

func main() {

	db, _ = gorm.Open("mysql", "dev:dev123@tcp(127.0.0.1:3306)/gotest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&Person{})

	r := gin.Default()
	r.GET("/", Home)
	r.GET("/person/", GetPersons)
	r.GET("/person/:id", GetPersonID)
	r.POST("/person", CreatePerson)
	r.PUT("/person/:id", UpdatePerson)
	r.DELETE("/person/:id", DeletePerson)
	r.Run(":8080")
}

// Home
func Home(c *gin.Context) {
	c.String(200, `Hello World
-------------------------------------		
GET    /person/                  
GET    /person/:id               
POST   /person                   
PUT    /person/:id               
DELETE /person/:id 
-------------------------------------
`)
}

// Getpersons all people from database (sqlite)
func GetPersons(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}

// GetPersonID data people by id
func GetPersonID(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person

	if err := db.Where("id= ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

// CreatePerson create new person
func CreatePerson(c *gin.Context) {
	var person Person
	c.BindJSON(&person)

	db.Create(&person)
	c.JSON(200, person)
}

// UpdatePerson update single people
func UpdatePerson(c *gin.Context) {
	var person Person
	id := c.Params.ByName("id")

	if err := db.Where("id= ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.BindJSON(&person)
	db.Save(&person)
	c.JSON(200, person)

}

func DeletePerson(c *gin.Context) {
	var person Person
	id := c.Params.ByName("id")

	d := db.Where("id= ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
