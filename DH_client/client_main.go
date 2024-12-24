package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting connection...")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	client := NewClient(conn, 128)
	client.keyAgreementClient()
	fmt.Println("Agreed key is ", client.key.String())
}
