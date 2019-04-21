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
	offer := webrtc.RTCSessionDescription{}
	if err != nil {
		// offer がない
		offer, err = peerConnection.CreateOffer(nil)
		c.SendOffer(ctx, &pb.OfferMessage{Token: token, Body: &pb.Offer{Body: offer.Sdp}})
		util.Check(err)
		fmt.Println(offer)
		// wait answer
		fmt.Println("waiting answer.")
		for {
			a, err := c.GetAnswer(ctx, token)
			if err != nil {
				time.Sleep(1 * time.Second)
				fmt.Print(".")
			} else {
				answer := webrtc.RTCSessionDescription{
					Type: webrtc.RTCSdpTypeAnswer,
					Sdp:  a.Body,
				}
				fmt.Println(answer)
				fmt.Println("answer was found")
				break
			}
		}
	} else {
		// offer があった
		offer = webrtc.RTCSessionDescription{
			Type: webrtc.RTCSdpTypeOffer,
			Sdp:  o.Body,
		}
		fmt.Println(offer)
		// send answer
		answer, err := peerConnection.CreateOffer(nil)
		_, err = c.SendAnswer(ctx, &pb.AnswerMessage{Token: token, Body: &pb.Answer{Body: answer.Sdp}})
		if err != nil {
			panic(err)
		}
		fmt.Println("-----------offer---------")
		fmt.Println(offer)
		fmt.Println("-----------answer---------")
		fmt.Println(answer)
	}
}
