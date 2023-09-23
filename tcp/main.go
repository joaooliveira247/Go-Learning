package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func server(ln net.Listener) {
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {

	defer c.Close()
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received:", msg)
}


func client(msg string) {
	c, err := net.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	fmt.Println("Sending...", msg)

	err = gob.NewEncoder(c).Encode(msg)

	if err != nil {
		fmt.Println(err)
	}
}

func readMsg() {
	for {
		input := bufio.NewReader(os.Stdin)
		msg, err := input.ReadString('\n')

		if msg == "" || err != nil {
			return
		}
		client(msg)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	go server(ln)
	readMsg()
}