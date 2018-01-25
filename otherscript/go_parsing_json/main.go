package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Username string
	Password string
	Domain   string
}

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	var conf *configuration

	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(conf.Username)
	fmt.Println(conf.Password)
	fmt.Println(conf.Domain)
}
