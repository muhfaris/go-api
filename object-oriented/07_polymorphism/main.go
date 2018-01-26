package main

import (
	"fmt"
)

type Human interface {
	myStereotype() string
}

type Man struct{}
type Woman struct{}

func (m Man) myStereotype() string {
	return "i'm going fishing"
}

func (w Woman) myStereotype() string {
	return "i'm going shopping"
}

func main() {
	m := new(Man)
	w := new(Woman)

	// an array of humans
	hArr := [...]Human{m, w}
	for n, _ := range hArr {
		fmt.Println("I human and my stereotype is:", hArr[n].myStereotype())
	}
}
