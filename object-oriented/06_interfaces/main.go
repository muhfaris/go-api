package main

/**
  In Go, the convention is to "er" a type to indicate that it is an interface. Shape is best named Shaper when it is an interface. Convert is best named Converter when it is an interface. Not doing so won’t crash your program but it is useful to follow convention as it will be easier for others reading the code to understand what is intended.
*/

import (
	"fmt"
)

// shaper interface
type Shaper interface {
	Area() int
}

// define a Rectangle struct has a length and widht
type Rectangle struct {
	length, width int
}

type Square struct {
	side int
}

// method
func (sq Square) Area() int {
	return sq.side * sq.side
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := Rectangle{length: 5, width: 3}
	fmt.Println("Details of rectangle:", r)

	var s Shaper
	s = r
	fmt.Println("Area of shaper (rectangle) :", s.Area())

	q := Square{5}
	fmt.Println("Square details are :", q)

	s = q
	fmt.Println("Area of shaper (Square)", s.Area())
	fmt.Println()

}
