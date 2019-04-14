package sdp

import (
	"context"
	"fmt"
	pb "github.com/takutakahashi/wg-connect/pkg/proto/sdp_exchange"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

func c(err error) {
	if err != nil {
		panic(err)
	}
	return
}

func (*server) GetPeer(ctx context.Context, in *pb.PeerMessage) (*pb.PeerResponse, error) {
	out := new(pb.PeerResponse)
	return out, nil
}

func (*server) GetAnswer(ctx context.Context, in *pb.Offer) (*pb.Answer, error) {
	out := new(pb.Answer)
	return out, nil
}

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
	c(err)
	fmt.Println(tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	c(err)
	defer listener.Close()
	s := grpc.NewServer()
	pb.RegisterExchangeServer(s, &server{})
	err = s.Serve(listener)
	c(err)
}
