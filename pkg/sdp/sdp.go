package sdp

import (
	"fmt"
	"net"
)

func echo(conn *net.TCPConn) {
	defer conn.Close()
	fmt.Println(conn.LocalAddr())
	for {
		buf := make([]byte, 1024*4)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func StartServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":50000")
	if err != nil {
		panic(err)
	}
	fmt.Println(tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go echo(conn)
	}
}
