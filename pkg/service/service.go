package service

import (
  pb "github.com/takutakahashi/wg-connect/pkg/proto/dist/sdp_exchange"
)

type GetPeerService struct {}

func (s *GetPeerService) GetPeer(message *pb.PeerMessage) (*pb.PeerResponse, error) {
}
