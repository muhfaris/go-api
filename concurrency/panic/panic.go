package main

import (
	"os"
)

func main() {
	var user = os.Getenv("icehiro")
	if user == "" {
		panic("No value for $USER")
	}
}
