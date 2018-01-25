package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "world", "a name say hello")
var english bool

func init() {
	flag.BoolVar(&english, "english", false, "Use english langugae")
	flag.BoolVar(&english, "s", false, "use english")
}

func main() {
	flag.Parse()

	if english == true {
		fmt.Printf("Holla %s\n", *name)
	} else {
		fmt.Printf("Hello %s\n", *name)
	}

}
