package main

import "fmt"

type Car struct {
	wheelCount int
}

type Ferarri struct {
	// anonymouse field car
	Car
}

type AstonMartin struct {
	Car
}

// define a behavior for Car
func (c Car) numberOfWhells() int {
	return c.wheelCount
}

func (f Ferarri) sayHi() {
	fmt.Println("Hi, Ferarri")
}

func (a AstonMartin) sayAston() {
	fmt.Println("Hi, Aston Martin")
}

func main() {
	f := Ferarri{Car{4}}
	fmt.Println("A ferarri has this many wheels:", f.numberOfWhells())
	f.sayHi()

	a := AstonMartin{Car{10}}
	fmt.Println("An Astom Martin has this many wheels:", a.numberOfWhells())
	a.sayAston()
}
