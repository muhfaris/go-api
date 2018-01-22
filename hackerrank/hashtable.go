package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	number := make([]int, 2)

	for count := 0; count < 2; count++ {
		fmt.Fscan(scanner, &number[count])
	}

	input := bufio.NewScanner(os.Stdin)
	string := make(map[int]int)

	for count := 0; count < 2; count++ {
		ok := input.Scan()
		if ok {
			lenword := WordCount(input.Text())
			string[count] = lenword
		}
	}
	if number[0] == string[0] && number[1] == string[1] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func WordCount(s string) int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return len(m)
}
