package main

import (
	"fmt"
)

type Rectangle struct {
	lenght, width int
}

func (r Rectangle) Area() int {
	return r.lenght * r.width
}

func (r *Rectangle) AreaReference() int {
	return r.lenght * r.width
}
func main() {
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is:", r1)
	fmt.Println("Rectangle area is:", r1.Area())
	fmt.Println("Rectangle by reference:", r1.AreaReference())
	fmt.Println("Rectangle area pointer:", (&r1).Area())
	fmt.Println("Rectangle area pointer by reference:", (&r1).AreaReference())

}
