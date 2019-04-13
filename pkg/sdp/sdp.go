package sdp

import (
	"fmt"
	"net"
  "google.golang.org/grpc"
  pb "github.com/takutakahashi/wg-connect/pkg/proto/sdp_exchange"
)

type server struct {}

func (s *server) GetPeer(message *pb.PeerMessage) (*pb.PeerResponse, error) {
  return &pb.PeerResponse{BodyJson: "{}"}, nil
}

func c(err error){
  if err != nil {
   panic(err)
  }
  return
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
}
