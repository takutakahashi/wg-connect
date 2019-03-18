package server

import (
	"fmt"
	"net"
)

func Listen() {
  addr := "localhost:9876"
  fmt.Println(addr)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}
