package model

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func handleChatClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("message to send")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		//send to socket
		fmt.Fprintf(conn, text+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(".", message)
	}
}
