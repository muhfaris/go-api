package main

import (
	"fmt"
	"time"
)

// Parameter type channel, type value is integer
func PrintCount(c chan int) {
	num := 0
	/* function to retrive data from c chanel
	   insert into variabel num
	*/
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	// create channel
	c := make(chan int)
	a := []int{2, 3, 4, 5, 6, 7, 8, 9, 0}

	// start go routin to PrintCount()
	go PrintCount(c)

	// Passes int into channel
	for _, v := range a {
		c <- v
	}

	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
