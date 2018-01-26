package main

import (
	"fmt"
)

type Kitchen struct {
	numOfLamps int
}

type House struct {
	Kitchen
	numOfLamps int
}

func main() {
	// Kitchen have 5 lamps
	// House have 10 lamps
	h := House{Kitchen{5}, 10}

	fmt.Println("House h has this many lamps:", h.numOfLamps)
	fmt.Println("The Kitchen in house h has this many lamps:", h.Kitchen.numOfLamps)
}
