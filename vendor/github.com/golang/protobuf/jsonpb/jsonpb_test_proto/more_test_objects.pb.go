// Code generated by protoc-gen-go.
// source: more_test_objects.proto
// DO NOT EDIT!

/*
Package jsonpb is a generated protocol buffer package.

It is generated from these files:
	more_test_objects.proto
	test_objects.proto

It has these top-level messages:
	Simple3
	Mappy
	Simple
	Repeats
	Widget
	Maps
	MsgWithOneof
	Real
	Complex
	KnownTypes
*/
package jsonpb

import "github.com/golang/protobuf/proto"
import "fmt"
import "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Numeral int32

const (
	Numeral_UNKNOWN Numeral = 0
	Numeral_ARABIC  Numeral = 1
	Numeral_ROMAN   Numeral = 2
)

var Numeral_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARABIC",
	2: "ROMAN",
}
var Numeral_value = map[string]int32{
	"UNKNOWN": 0,
	"ARABIC":  1,
	"ROMAN":   2,
}

func (x Numeral) String() string {
	return proto.EnumName(Numeral_name, int32(x))
}
func (Numeral) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Simple3 struct {
	Dub float64 `protobuf:"fixed64,1,opt,name=dub" json:"dub,omitempty"`
}

func (m *Simple3) Reset()                    { *m = Simple3{} }
func (m *Simple3) String() string            { return proto.CompactTextString(m) }
func (*Simple3) ProtoMessage()               {}
func (*Simple3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Mappy struct {
	Nummy    map[int64]int32    `protobuf:"bytes,1,rep,name=nummy" json:"nummy,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Strry    map[string]string  `protobuf:"bytes,2,rep,name=strry" json:"strry,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Objjy    map[int32]*Simple3 `protobuf:"bytes,3,rep,name=objjy" json:"objjy,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Buggy    map[int64]string   `protobuf:"bytes,4,rep,name=buggy" json:"buggy,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Booly    map[bool]bool      `protobuf:"bytes,5,rep,name=booly" json:"booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Enumy    map[string]Numeral `protobuf:"bytes,6,rep,name=enumy" json:"enumy,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=jsonpb.Numeral"`
	S32Booly map[int32]bool     `protobuf:"bytes,7,rep,name=s32booly" json:"s32booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	S64Booly map[int64]bool     `protobuf:"bytes,8,rep,name=s64booly" json:"s64booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	U32Booly map[uint32]bool    `protobuf:"bytes,9,rep,name=u32booly" json:"u32booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	U64Booly map[uint64]bool    `protobuf:"bytes,10,rep,name=u64booly" json:"u64booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *Mappy) Reset()                    { *m = Mappy{} }
func (m *Mappy) String() string            { return proto.CompactTextString(m) }
func (*Mappy) ProtoMessage()               {}
func (*Mappy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Mappy) GetNummy() map[int64]int32 {
	if m != nil {
		return m.Nummy
	}
	return nil
}

func (m *Mappy) GetStrry() map[string]string {
	if m != nil {
		return m.Strry
	}
	return nil
}

func (m *Mappy) GetObjjy() map[int32]*Simple3 {
	if m != nil {
		return m.Objjy
	}
	return nil
}

func (m *Mappy) GetBuggy() map[int64]string {
	if m != nil {
		return m.Buggy
	}
	return nil
}

func (m *Mappy) GetBooly() map[bool]bool {
	if m != nil {
		return m.Booly
	}
	return nil
}

func (m *Mappy) GetEnumy() map[string]Numeral {
	if m != nil {
		return m.Enumy
	}
	return nil
}

func (m *Mappy) GetS32Booly() map[int32]bool {
	if m != nil {
		return m.S32Booly
	}
	return nil
}

func (m *Mappy) GetS64Booly() map[int64]bool {
	if m != nil {
		return m.S64Booly
	}
	return nil
}

func (m *Mappy) GetU32Booly() map[uint32]bool {
	if m != nil {
		return m.U32Booly
	}
	return nil
}

func (m *Mappy) GetU64Booly() map[uint64]bool {
	if m != nil {
		return m.U64Booly
	}
	return nil
}

func init() {
	proto.RegisterType((*Simple3)(nil), "jsonpb.Simple3")
	proto.RegisterType((*Mappy)(nil), "jsonpb.Mappy")
	proto.RegisterEnum("jsonpb.Numeral", Numeral_name, Numeral_value)
}

func init() { proto.RegisterFile("more_test_objects.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6b, 0xdb, 0x30,
	0x14, 0x87, 0xe7, 0xa4, 0x4e, 0xec, 0x17, 0xba, 0x19, 0x31, 0x98, 0x58, 0x2f, 0xa1, 0x30, 0x08,
	0x83, 0xf9, 0x90, 0x8c, 0xad, 0x6c, 0xa7, 0x74, 0xf4, 0x50, 0x46, 0x1d, 0x70, 0x09, 0x3b, 0x96,
	0x78, 0x13, 0x65, 0x9e, 0x6d, 0x19, 0xdb, 0x1a, 0xe8, 0x8f, 0x1f, 0x8c, 0x27, 0xcb, 0xb5, 0x6c,
	0x14, 0xd2, 0x9b, 0xcc, 0xef, 0xfb, 0xf2, 0x9e, 0xf4, 0x1e, 0x81, 0x37, 0x39, 0xaf, 0xd8, 0x43,
	0xc3, 0xea, 0xe6, 0x81, 0x27, 0x29, 0xfb, 0xd9, 0xd4, 0x61, 0x59, 0xf1, 0x86, 0x93, 0x59, 0x5a,
	0xf3, 0xa2, 0x4c, 0x2e, 0x2f, 0x60, 0x7e, 0xff, 0x3b, 0x2f, 0x33, 0xb6, 0x21, 0x01, 0x4c, 0x7f,
	0x89, 0x84, 0x3a, 0x4b, 0x67, 0xe5, 0xc4, 0x78, 0xbc, 0xfc, 0xe7, 0x81, 0x7b, 0x77, 0x28, 0x4b,
	0x49, 0x42, 0x70, 0x0b, 0x91, 0xe7, 0x92, 0x3a, 0xcb, 0xe9, 0x6a, 0xb1, 0xa6, 0x61, 0xab, 0x87,
	0x2a, 0x0d, 0x23, 0x8c, 0x6e, 0x8a, 0xa6, 0x92, 0x71, 0x8b, 0x21, 0x5f, 0x37, 0x55, 0x25, 0xe9,
	0xc4, 0xc6, 0xdf, 0x63, 0xa4, 0x79, 0x85, 0x21, 0xcf, 0x93, 0x34, 0x95, 0x74, 0x6a, 0xe3, 0x77,
	0x18, 0x69, 0x5e, 0x61, 0xc8, 0x27, 0xe2, 0xf1, 0x51, 0xd2, 0x33, 0x1b, 0x7f, 0x8d, 0x91, 0xe6,
	0x15, 0xa6, 0x78, 0xce, 0x33, 0x49, 0x5d, 0x2b, 0x8f, 0x51, 0xc7, 0xe3, 0x19, 0x79, 0x56, 0x88,
	0x5c, 0xd2, 0x99, 0x8d, 0xbf, 0xc1, 0x48, 0xf3, 0x0a, 0x23, 0x9f, 0xc1, 0xab, 0x37, 0xeb, 0xb6,
	0xc4, 0x5c, 0x29, 0x17, 0xa3, 0x2b, 0xeb, 0xb4, 0xb5, 0x9e, 0x60, 0x25, 0x7e, 0xfa, 0xd8, 0x8a,
	0x9e, 0x55, 0xd4, 0x69, 0x27, 0xea, 0x4f, 0x14, 0x45, 0x57, 0xd1, 0xb7, 0x89, 0xfb, 0x61, 0x45,
	0x61, 0x54, 0x14, 0x5d, 0x45, 0xb0, 0x8a, 0xc3, 0x8a, 0x1d, 0xfc, 0xf6, 0x0a, 0xa0, 0x1f, 0x34,
	0x6e, 0xcb, 0x1f, 0x26, 0xd5, 0xb6, 0x4c, 0x63, 0x3c, 0x92, 0xd7, 0xe0, 0xfe, 0x3d, 0x64, 0x82,
	0xd1, 0xc9, 0xd2, 0x59, 0xb9, 0x71, 0xfb, 0xf1, 0x65, 0x72, 0xe5, 0xa0, 0xd9, 0x8f, 0xdc, 0x34,
	0x7d, 0x8b, 0xe9, 0x9b, 0xe6, 0x2d, 0x40, 0x3f, 0x7c, 0xd3, 0x74, 0x5b, 0xf3, 0x9d, 0x69, 0x2e,
	0xd6, 0xaf, 0xba, 0x9b, 0xe8, 0x9d, 0x1e, 0x35, 0xd1, 0xef, 0xc5, 0xa9, 0xf6, 0xfd, 0xb1, 0xf9,
	0xf4, 0x20, 0xa6, 0xe9, 0x59, 0x4c, 0x6f, 0xd4, 0x7e, 0xbf, 0x2b, 0x96, 0x8b, 0x0f, 0xda, 0x7f,
	0xd9, 0xb7, 0x1f, 0x89, 0x9c, 0x55, 0x87, 0xcc, 0xfc, 0xa9, 0xaf, 0x70, 0x3e, 0xd8, 0x21, 0xcb,
	0x63, 0x1c, 0xef, 0x03, 0x65, 0x73, 0xaa, 0xa7, 0xae, 0x3f, 0x96, 0xf7, 0xc7, 0x2a, 0x9f, 0x3f,
	0x47, 0x3e, 0x56, 0xf9, 0xec, 0x84, 0xfc, 0xfe, 0x03, 0xcc, 0xf5, 0x4b, 0x90, 0x05, 0xcc, 0xf7,
	0xd1, 0xf7, 0x68, 0xf7, 0x23, 0x0a, 0x5e, 0x10, 0x80, 0xd9, 0x36, 0xde, 0x5e, 0xdf, 0x7e, 0x0b,
	0x1c, 0xe2, 0x83, 0x1b, 0xef, 0xee, 0xb6, 0x51, 0x30, 0x49, 0x66, 0xea, 0xaf, 0x6d, 0xf3, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0xa2, 0x4b, 0xe1, 0x77, 0xf5, 0x04, 0x00, 0x00,
}
