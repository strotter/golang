package main

import (
	"fmt"
	"net"
	"encoding/gob"
)

type P struct {
	M, N int64
}

func handleConn(conn net.Conn) {
	fmt.Printf("Got new connection.")
	dec := gob.NewDecoder(conn)
	p := &P{}
	dec.Decode(p)
	fmt.Printf("Received: %+v", p);
}

func main() {
	fmt.Println("Server start")
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error on accept: ", err)
			continue
		}

		go handleConn(conn)
	}
}