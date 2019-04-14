// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sdp_exchange.proto

package sdp_exchange

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PeerMessage struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerMessage) Reset()         { *m = PeerMessage{} }
func (m *PeerMessage) String() string { return proto.CompactTextString(m) }
func (*PeerMessage) ProtoMessage()    {}
func (*PeerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{0}
}

func (m *PeerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerMessage.Unmarshal(m, b)
}
func (m *PeerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerMessage.Marshal(b, m, deterministic)
}
func (m *PeerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerMessage.Merge(m, src)
}
func (m *PeerMessage) XXX_Size() int {
	return xxx_messageInfo_PeerMessage.Size(m)
}
func (m *PeerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PeerMessage proto.InternalMessageInfo

func (m *PeerMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PeerResponse struct {
	BodyJson             string   `protobuf:"bytes,2,opt,name=body_json,json=bodyJson,proto3" json:"body_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerResponse) Reset()         { *m = PeerResponse{} }
func (m *PeerResponse) String() string { return proto.CompactTextString(m) }
func (*PeerResponse) ProtoMessage()    {}
func (*PeerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{1}
}

func (m *PeerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerResponse.Unmarshal(m, b)
}
func (m *PeerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerResponse.Marshal(b, m, deterministic)
}
func (m *PeerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerResponse.Merge(m, src)
}
func (m *PeerResponse) XXX_Size() int {
	return xxx_messageInfo_PeerResponse.Size(m)
}
func (m *PeerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeerResponse proto.InternalMessageInfo

func (m *PeerResponse) GetBodyJson() string {
	if m != nil {
		return m.BodyJson
	}
	return ""
}

type Offer struct {
	SdpBody              string   `protobuf:"bytes,3,opt,name=sdp_body,json=sdpBody,proto3" json:"sdp_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Offer) Reset()         { *m = Offer{} }
func (m *Offer) String() string { return proto.CompactTextString(m) }
func (*Offer) ProtoMessage()    {}
func (*Offer) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{2}
}

func (m *Offer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Offer.Unmarshal(m, b)
}
func (m *Offer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Offer.Marshal(b, m, deterministic)
}
func (m *Offer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Offer.Merge(m, src)
}
func (m *Offer) XXX_Size() int {
	return xxx_messageInfo_Offer.Size(m)
}
func (m *Offer) XXX_DiscardUnknown() {
	xxx_messageInfo_Offer.DiscardUnknown(m)
}

var xxx_messageInfo_Offer proto.InternalMessageInfo

func (m *Offer) GetSdpBody() string {
	if m != nil {
		return m.SdpBody
	}
	return ""
}

type Answer struct {
	SdpBody              string   `protobuf:"bytes,4,opt,name=sdp_body,json=sdpBody,proto3" json:"sdp_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{3}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetSdpBody() string {
	if m != nil {
		return m.SdpBody
	}
	return ""
}

func init() {
	proto.RegisterType((*PeerMessage)(nil), "PeerMessage")
	proto.RegisterType((*PeerResponse)(nil), "PeerResponse")
	proto.RegisterType((*Offer)(nil), "Offer")
	proto.RegisterType((*Answer)(nil), "Answer")
}

func init() { proto.RegisterFile("sdp_exchange.proto", fileDescriptor_c88d59dbcf37a6c3) }

var fileDescriptor_c88d59dbcf37a6c3 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0xef, 0x55, 0x6f, 0x7f, 0xc6, 0xba, 0x09, 0x2e, 0x6a, 0x05, 0x91, 0x74, 0x53, 0x10,
	0xb2, 0xd0, 0x27, 0x50, 0x90, 0x82, 0x20, 0x4a, 0x71, 0x5f, 0x5a, 0x33, 0xad, 0x28, 0x64, 0x42,
	0x27, 0xa0, 0x7d, 0x7b, 0x49, 0xda, 0x45, 0x75, 0x39, 0x67, 0x3e, 0x66, 0xbe, 0x03, 0x82, 0xb5,
	0x6d, 0xf1, 0xe7, 0xfd, 0xa3, 0x33, 0x23, 0x2a, 0x3b, 0x91, 0x23, 0x59, 0xc2, 0xe9, 0x2b, 0xe2,
	0xf4, 0x8c, 0xcc, 0xdd, 0x88, 0xe2, 0x1c, 0x0e, 0x8e, 0xbe, 0xd0, 0xe4, 0xfb, 0xeb, 0x7d, 0x95,
	0x36, 0xcb, 0x20, 0x6f, 0x20, 0xf3, 0x50, 0x83, 0x6c, 0xc9, 0x30, 0x8a, 0x4b, 0x48, 0x7b, 0xd2,
	0x73, 0xfb, 0xc9, 0x64, 0xf2, 0xa3, 0x40, 0x26, 0x3e, 0x78, 0x62, 0x32, 0x52, 0xc2, 0xe1, 0x65,
	0x18, 0x70, 0x12, 0x17, 0x90, 0xf8, 0x87, 0x7e, 0x91, 0x1f, 0x07, 0x28, 0x66, 0x6d, 0x1f, 0x48,
	0xcf, 0xb2, 0x84, 0xe8, 0xde, 0xf0, 0xf7, 0x3f, 0xe8, 0xe4, 0x0f, 0x74, 0xfb, 0x06, 0xc9, 0xe3,
	0x2a, 0x2b, 0x2a, 0x88, 0x6b, 0x74, 0x5e, 0x42, 0x64, 0x6a, 0x23, 0x5c, 0x9c, 0xa9, 0xad, 0x99,
	0xdc, 0x89, 0x2b, 0x48, 0x6b, 0x74, 0xeb, 0xf5, 0x48, 0x05, 0x95, 0x22, 0x56, 0x4b, 0x20, 0x77,
	0x7d, 0x14, 0x7a, 0xdf, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x23, 0x27, 0x7a, 0x0d, 0x01,
	0x00, 0x00,
}