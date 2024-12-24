package main

import (
	"fmt"
	"math/big"
	"net"
)

type dh_client struct {
	connection net.Conn // Go's version of Socket
	key_size   int
	g          big.Int
	p          big.Int
	a          big.Int
	A          big.Int
	B          big.Int
	key        big.Int
}

func NewClient(con net.Conn, size int) *dh_client {
	return &dh_client{
		connection: con,
		key_size:   size,
		g:          *primeGen(size),
		p:          *new(big.Int),
		a:          *primeGen(size),
		A:          *new(big.Int),
		B:          *new(big.Int),
		key:        *new(big.Int),
	}
}

func (client *dh_client) keyAgreementClient() {
	// Diffie-Hellman algorithm for key agreement. This is for the client side.
	conn := client.connection
	_, err := conn.Write(client.g.Bytes()) //send g value to server
	if err != nil {
		fmt.Println(err)
		return
	}

	byte_len := (client.key_size + 7) / 8
	buf := make([]byte, byte_len)
	_, err2 := conn.Read(buf) //receive p value from server
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	client.p.SetBytes(buf)                        //set p to the received p value
	client.A.Exp(&client.g, &client.a, &client.p) //calculate A

	_, err3 := conn.Write(client.A.Bytes()) //send A value to server
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	buf = make([]byte, byte_len)
	_, err4 := conn.Read(buf) //receive B value from server
	if err4 != nil {
		fmt.Println(err4)
		return
	}
	client.B.SetBytes(buf) //set B to the received B value
	client.key.Exp(&client.B, &client.a, &client.p)
}
