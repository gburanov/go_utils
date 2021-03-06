// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package goutils is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
*/
package goutils

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	DriverId         *int64                     `protobuf:"varint,1,req,name=driverId" json:"driverId,omitempty"`
	Latitude         *float64                   `protobuf:"fixed64,2,req,name=latitude" json:"latitude,omitempty"`
	Longitude        *float64                   `protobuf:"fixed64,3,req,name=longitude" json:"longitude,omitempty"`
	Stamp            *google_protobuf.Timestamp `protobuf:"bytes,4,req,name=stamp" json:"stamp,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetDriverId() int64 {
	if m != nil && m.DriverId != nil {
		return *m.DriverId
	}
	return 0
}

func (m *Message) GetLatitude() float64 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *Message) GetLongitude() float64 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *Message) GetStamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Stamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "goutils.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xcf, 0x2f, 0x2d, 0xc9,
	0xcc, 0x29, 0x96, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x27, 0x95, 0xa6,
	0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x54, 0x2a, 0x4d, 0x64, 0xe4,
	0x62, 0xf7, 0x85, 0xe8, 0x15, 0x92, 0xe2, 0xe2, 0x48, 0x29, 0xca, 0x2c, 0x4b, 0x2d, 0xf2, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd2, 0x60, 0x0e, 0x82, 0xf3, 0x41, 0x72, 0x39, 0x89, 0x25, 0x99, 0x25,
	0xa5, 0x29, 0xa9, 0x12, 0x4c, 0x0a, 0x4c, 0x1a, 0x8c, 0x41, 0x70, 0xbe, 0x90, 0x0c, 0x17, 0x67,
	0x4e, 0x7e, 0x5e, 0x3a, 0x44, 0x92, 0x19, 0x2c, 0x89, 0x10, 0x10, 0x32, 0xe0, 0x62, 0x05, 0x5b,
	0x28, 0xc1, 0xa2, 0xc0, 0xa4, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0x71, 0x92, 0x1e, 0xcc, 0x49, 0x7a,
	0x21, 0x30, 0x27, 0x05, 0x41, 0x14, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x81, 0xf5, 0x02,
	0xcd, 0x00, 0x00, 0x00,
}
