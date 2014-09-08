package main

import (
	"fmt"
	//"math/rand"
	"net"
	"bufio"
	"strconv"
)

const PORT = 3333

func main() {
	server, err := net.Listen("tcp", ":" + strconv.Itoa(PORT))

	if server == nil {
		panic("Couldn't start listening: " + err.Error())
	}

	conns := clientConns(server)

	for {
		go handleConn(<-conns)
	}
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0

	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Println("Couldn't accept: " + err.Error())
				continue
			}

			i++
			fmt.Printf("%d: %v <-> %v", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()

	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)

	for {
		line, err := b.ReadString('\n')

		if err != nil {
			fmt.Println("Error on client: ", err.Error())
			break
		}

		fmt.Println("Received: ", line)
		client.Write([]byte(line))
	}
}