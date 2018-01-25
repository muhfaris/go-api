package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://golang.org")
	fmt.Println(resp)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	resp.Body.Close()
}
