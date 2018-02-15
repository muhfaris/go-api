package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Atoi convert string to int

	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	result := sum(a, b)
	fmt.Printf("The sum of %d and %d is %d", a, b, result)
}

func sum(a, b int) int {
	return a + b
}
