package sdp

import (
	"context"
	"errors"
	"fmt"
	"github.com/pions/webrtc"
	pb "github.com/takutakahashi/wg-connect/pkg/proto/sdp_exchange"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	ch             chan string
	description    chan webrtc.RTCSessionDescription
	offerPool      map[string]webrtc.RTCSessionDescription
	peerConnection *webrtc.RTCPeerConnection
}

func c(err error) {
	if err != nil {
		panic(err)
	}
	return
}

func (s *server) GetOffer(ctx context.Context, in *pb.Token) (*pb.Offer, error) {
	fmt.Println(in)
	if v, ok := s.offerPool[in.Body]; ok {
		out := &pb.Offer{Body: v.Sdp}
		return out, nil
	} else {
		return nil, errors.New("offer is not found")
	}
}

func (s *server) SendOffer(ctx context.Context, in *pb.OfferMessage) (*pb.OfferResponse, error) {
	s.offerPool[in.Token.Body] = webrtc.RTCSessionDescription{
		Type: webrtc.RTCSdpTypeOffer,
		Sdp:  in.Body,
	}
	fmt.Println(s.offerPool)
	return &pb.OfferResponse{Code: "ok"}, nil
}

func (s *server) GetPeer(ctx context.Context, in *pb.PeerMessage) (*pb.PeerResponse, error) {
	out := new(pb.PeerResponse)
	out.BodyJson = "{}"

	s.ch <- in.Token

	return out, nil
}

func (*server) GetAnswer(ctx context.Context, in *pb.Offer) (*pb.Answer, error) {
	out := new(pb.Answer)
	return out, nil
}

func StartServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":50000")
	c(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	c(err)
	defer listener.Close()
	config := webrtc.RTCConfiguration{
		IceServers: []webrtc.RTCIceServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	peerConnection, err := webrtc.New(config)
	s := grpc.NewServer()
	serve := &server{
		ch:             make(chan string),
		description:    make(chan webrtc.RTCSessionDescription),
		offerPool:      map[string]webrtc.RTCSessionDescription{},
		peerConnection: peerConnection,
	}
	pb.RegisterExchangeServer(s, serve)

	go func() {
		for mes := range serve.ch {
			fmt.Println(mes)
		}
	}()

	err = s.Serve(listener)
	c(err)
}
