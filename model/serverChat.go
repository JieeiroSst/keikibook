package model

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleChatServer() {
	fmt.Println("loading server....")

	//listen
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	//for vong lap xu ly event
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("message reciew", string(message))

		newMessage := strings.ToLower(message)

		conn.Write([]byte(newMessage + "\n"))
	}
}
