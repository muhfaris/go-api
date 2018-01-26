package main

import (
	"fmt"
)

type Retangle struct {
	lenght, width int
	name          string
}

func main() {

	// initialize value in order
	r := Retangle{10, 5, "value in order"}

	// initialize value by variabel
	r1 := Retangle{lenght: 14, width: 7, name: "value by variabel"}

	// get pointer to an instead new keyword
	pr := new(Retangle)
	pr.lenght = 16
	pr.width = 8
	pr.name = "value by pointer"

	fmt.Println("Default Retangle is :", r)
	fmt.Println("Default Retangle is :", r1)
	fmt.Println("Default Retangle is :", pr)

	/**
	Default Retangle is : {10 5 value in order}
	Default Retangle is : {14 7 value by variabel}
	Default Retangle is : &{16 8 value by pointer}
	*/
}

/**
Encapsulation and visibility of structs and variables

- variable starts with capital letter, so visible outside this package
- variable starts with small letter, so NOT visible outside package
- variable starts with capital letter, so visible outside package
- variabel not capital not exported
- variabel capital exported

type notExported struct {
}

type Exported struct {
    notExportedVariable int
    ExportedVariable int
    s string
    S string
}

**/
