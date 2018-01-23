package main

import (
	"fmt"
)

type Human struct {
	Name  string
	Age   int
	Phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Println("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}

func (h Human) Sing(lyric string) {
	fmt.Println("la ... la... la..la ...", lyric)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s , I work at %s. Call me on %s\n", e.Name, e.company, e.Phone)
}

//interface
type Men interface {
	SayHi()
	Sing(lyric string)
}

func main() {
	mike := Student{Human{"mike", 24, "23453234"}, "MIT", 0.00}
	sam := Student{Human{"sam", 20, "2342333323"}, "Golang Inc", 1000}

	//define ineterface i
	var i Men
	i = mike
	fmt.Println("This is mike, a student:")
	i.SayHi()
	i.Sing("Indonesia Tanah air...")

	fmt.Println("lets use a slice of Men and see what happenes")
	x := make([]Men, 2)

	x[0], x[1] = mike, sam

	for _, value := range x {
		value.SayHi()
	}
}
