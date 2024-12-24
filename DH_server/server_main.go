package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening for a connection...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	server := NewServer(conn, 128)
	server.keyAgreementServer()
	fmt.Println("Agreed key is ", server.key.String())
}
