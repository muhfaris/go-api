package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muhfaris/goFun/basic_apiseparate/models"
)

type (
	UserController struct{}
)

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Ali rahman",
		Gender: "Male",
		Age:    22,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)

	// Write content, status code
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser create new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// retrive post data alias u variabel
	u := models.User{}

	// decode data from json
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "foo"

	uj, _ := json.Marshal(u)

	// write content and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser remove user by id
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
}
