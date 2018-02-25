package main

import (
	"fmt"
	"net"
)

func main() {

	if listener, err := net.Listen("tcp", ":4040"); err != nil {
		fmt.Println(err)
		return
	} else {
		defer listener.Close()
		for {
			if conn, err := listener.Accept(); err != nil {
				fmt.Println(err)
				return
			} else {

				conn.Write([]byte("Welcome to listen go"))
				conn.Close()
			}

		}
	}

}
