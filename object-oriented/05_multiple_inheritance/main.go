package main

import (
	"fmt"
)

type Camera struct{}
type Phone struct{}

// multiple inheritance
type CameraPhone struct {
	//anonymous camera n phone
	Camera
	Phone
}

// not using the type, so discard it by putting
func (_ Camera) takePicture() string {
	return "Click"
}

// not using the type, so discard its by putting a _
func (_ Phone) call() string {
	return "Ring Ring"
}

func main() {
	// a new camera phone instance
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone multiple behavior ... ")
	fmt.Println("It can take picture:", cp.takePicture())
	fmt.Println("It can also make calls:", cp.call())

}
