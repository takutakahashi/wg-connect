package client

import (
	"fmt"
	"net"
)

func ConnectServer() {
	conn, err := net.Dial("udp", "localhost:9876")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("hello from client"))
	for {
		buffer := make([]byte, 1500)
		length, _ := conn.Read(buffer)
		fmt.Printf("Receive: %v\n", string(buffer[:length]))
	}

}
