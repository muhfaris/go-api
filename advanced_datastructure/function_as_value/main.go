package main

import (
	"fmt"
)

func IsOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	} else {
		return true
	}
}

func IsEven(integer int) bool {
	if integer%2 == 0 {
		return true
	} else {
		return false
	}

}

func filter(slice []int, f func(int) bool) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	odd := filter(slice, IsOdd)
	fmt.Println(odd)

	even := filter(slice, IsEven)
	fmt.Println(even)
}
