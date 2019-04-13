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

type SDPExchange struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SdpBody              string   `protobuf:"bytes,2,opt,name=sdp_body,json=sdpBody,proto3" json:"sdp_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SDPExchange) Reset()         { *m = SDPExchange{} }
func (m *SDPExchange) String() string { return proto.CompactTextString(m) }
func (*SDPExchange) ProtoMessage()    {}
func (*SDPExchange) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d59dbcf37a6c3, []int{0}
}

func (m *SDPExchange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SDPExchange.Unmarshal(m, b)
}
func (m *SDPExchange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SDPExchange.Marshal(b, m, deterministic)
}
func (m *SDPExchange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SDPExchange.Merge(m, src)
}
func (m *SDPExchange) XXX_Size() int {
	return xxx_messageInfo_SDPExchange.Size(m)
}
func (m *SDPExchange) XXX_DiscardUnknown() {
	xxx_messageInfo_SDPExchange.DiscardUnknown(m)
}

var xxx_messageInfo_SDPExchange proto.InternalMessageInfo

func (m *SDPExchange) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SDPExchange) GetSdpBody() string {
	if m != nil {
		return m.SdpBody
	}
	return ""
}

func init() {
	proto.RegisterType((*SDPExchange)(nil), "SDPExchange")
}

func init() { proto.RegisterFile("sdp_exchange.proto", fileDescriptor_c88d59dbcf37a6c3) }

var fileDescriptor_c88d59dbcf37a6c3 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0x29, 0x88,
	0x4f, 0xad, 0x48, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2,
	0xe3, 0xe2, 0x0e, 0x76, 0x09, 0x70, 0x85, 0x0a, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0xe4, 0x67, 0xa7,
	0xe6, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x92, 0x5c, 0x1c, 0x20, 0xad,
	0x49, 0xf9, 0x29, 0x95, 0x12, 0x4c, 0x60, 0x09, 0xf6, 0xe2, 0x94, 0x02, 0xa7, 0xfc, 0x94, 0xca,
	0x24, 0x36, 0xb0, 0x31, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0xf6, 0x50, 0xfa, 0x5c,
	0x00, 0x00, 0x00,
}