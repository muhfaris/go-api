package main

import (
	// standard library packakges

	"net/http"

	// third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/muhfaris/goFun/basic_apiseparate/controllers"
)

func main() {
	// inisialisasi
	r := httprouter.New()

	// Get UserController instead
	uc := controllers.NewUserController()

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.RemoveUser)
	// go fly
	http.ListenAndServe("localhost:8181", r)
}
