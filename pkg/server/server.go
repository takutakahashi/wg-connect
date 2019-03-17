package server

import (
	"fmt"
	"net"
)

func Listen() {
  addr := "localhost:9876"
  fmt.Println(addr)
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddr, _ := conn.ReadFrom(buffer)
		fmt.Printf("Received from %v: %v\n", remoteAddr, string(buffer[:length]))
		conn.WriteTo([]byte("hello"), remoteAddr)
	}
}
