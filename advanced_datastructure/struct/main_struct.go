package main

import (
	"fmt"
)

type Human struct {
	Name  string
	Age   int
	Phone string
}

type Employes struct {
	Human
	Speciality string
	Phone      string
}

func main() {
	Bob := Employes{Human{"Bob", 20, "234567876"}, "Designer", "12121212"}
	fmt.Println("Bob's work phone is:", Bob.Phone)

	fmt.Println("Phone from Human :", Bob.Human.Phone)
}
