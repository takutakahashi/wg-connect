package client

import (
	"context"
	"fmt"
	"github.com/pions/webrtc"
	"github.com/pions/webrtc/examples/util"
	pb "github.com/takutakahashi/wg-connect/pkg/proto/sdp_exchange"
	"google.golang.org/grpc"
	"time"
)

func ConnectServer() {
	config := webrtc.RTCConfiguration{
		IceServers: []webrtc.RTCIceServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	peerConnection, err := webrtc.New(config)
	util.Check(err)
	token := &pb.Token{Body: "takutaka"}
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
	util.Check(err)
	defer conn.Close()
	c := pb.NewExchangeClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	o, err := c.GetOffer(ctx, token)
	if err != nil {
		// offer がない
		fmt.Println(err)
		r, err := c.GetPeer(ctx, &pb.PeerMessage{Token: "takutaka"})
		offer, err := peerConnection.CreateOffer(nil)
		c.SendOffer(ctx, &pb.OfferMessage{Token: token, Body: offer.Sdp})
		util.Check(err)
		fmt.Println(r)
	} else {
		// offer があった
		offer := webrtc.RTCSessionDescription{
			Type: webrtc.RTCSdpTypeOffer,
			Sdp:  o.Body,
		}
		fmt.Println(offer)
		fmt.Println("offer found")
	}
}
