package sdp_exchange

import (
	"context"
	"google.golang.org/grpc"
)

type ExchangeClient interface {
	// Sends a greeting
	GetPeer(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerResponse, error)
}

type exchangeClient struct {
	cc *grpc.ClientConn
}

func NewExchangeClient(cc *grpc.ClientConn) ExchangeClient {
	return &exchangeClient{cc}
}

func GetPeer(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerResponse, error) {
	res := new(PeerResponse)
	return res, nil
}
