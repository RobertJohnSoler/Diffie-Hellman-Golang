package main

import (
	"Diffie-Hellman-Golang/utils"
	"fmt"
	"math/big"
	"net"
)

type dh_server struct {
	connection net.Conn // Go's version of Socket
	key_size   int
	g          big.Int
	p          big.Int
	b          big.Int
	B          big.Int
	A          big.Int
	key        big.Int
}

func NewServer(con net.Conn, size int) *dh_server {
	return &dh_server{
		connection: con,
		key_size:   size,
		g:          *new(big.Int),
		p:          *utils.PrimeGen(size),
		b:          *utils.PrimeGen(size),
		B:          *new(big.Int),
		A:          *new(big.Int),
		key:        *new(big.Int),
	}
}

func (server *dh_server) keyAgreementServer() {
	// Diffie-Hellman algorithm for key agreement. This is for the server side.
	conn := server.connection
	byte_len := (server.key_size + 7) / 8
	buf := make([]byte, byte_len)
	_, err := conn.Read(buf) //receive g value from client
	if err != nil {
		fmt.Println(err)
		return
	}
	server.g.SetBytes(buf) //set g to the received g value

	_, err2 := conn.Write(server.p.Bytes()) //send p value to client
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	buf = make([]byte, byte_len)
	_, err4 := conn.Read(buf) //receive A value from client
	if err4 != nil {
		fmt.Println(err)
		return
	}
	server.A.SetBytes(buf) //set A to the received A value

	server.B.Exp(&server.g, &server.b, &server.p) //calculate B

	_, err3 := conn.Write(server.B.Bytes()) //send B value to client
	if err3 != nil {
		fmt.Println(err)
		return
	}

	server.key.Exp(&server.A, &server.b, &server.p)

}
