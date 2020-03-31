// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filter_message.proto

package pb

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

type Filter struct {
	MaxPriceUsd          float64  `protobuf:"fixed64,1,opt,name=max_price_usd,json=maxPriceUsd,proto3" json:"max_price_usd,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,2,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MinCpuCore           uint32   `protobuf:"varint,3,opt,name=min_cpu_core,json=minCpuCore,proto3" json:"min_cpu_core,omitempty"`
	MinRam               *Memory  `protobuf:"bytes,4,opt,name=min_ram,json=minRam,proto3" json:"min_ram,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_02dd12c5821a9fa1, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetMaxPriceUsd() float64 {
	if m != nil {
		return m.MaxPriceUsd
	}
	return 0
}

func (m *Filter) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *Filter) GetMinCpuCore() uint32 {
	if m != nil {
		return m.MinCpuCore
	}
	return 0
}

func (m *Filter) GetMinRam() *Memory {
	if m != nil {
		return m.MinRam
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "laptopstore.pcbook.Filter")
}

func init() {
	proto.RegisterFile("filter_message.proto", fileDescriptor_02dd12c5821a9fa1)
}

var fileDescriptor_02dd12c5821a9fa1 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x49, 0x2d, 0x2b, 0xa4, 0xf6, 0x12, 0x0a, 0x2e, 0x3d, 0x85, 0x9e, 0xf6, 0x94, 0x83,
	0x7d, 0x03, 0x0b, 0x7a, 0x12, 0x64, 0xc1, 0x8b, 0x97, 0x90, 0x4d, 0xc7, 0x36, 0xd8, 0xd9, 0x19,
	0x26, 0x59, 0xa8, 0x7d, 0x16, 0x1f, 0x56, 0x1a, 0x4f, 0xe2, 0x75, 0xbe, 0xf9, 0x3f, 0x3e, 0xbd,
	0xfa, 0x48, 0xa7, 0x02, 0xe2, 0x11, 0x72, 0x0e, 0x07, 0x70, 0x2c, 0x54, 0xc8, 0x98, 0x53, 0xe0,
	0x42, 0x9c, 0x0b, 0x09, 0x38, 0x8e, 0x03, 0xd1, 0xe7, 0x7a, 0x85, 0x80, 0x24, 0x5f, 0x7f, 0x3f,
	0x37, 0xdf, 0x4a, 0x37, 0x4f, 0x55, 0x61, 0x36, 0x7a, 0x89, 0xe1, 0xec, 0x59, 0x52, 0x04, 0x3f,
	0xe5, 0x7d, 0xab, 0xac, 0xea, 0x54, 0xbf, 0xc0, 0x70, 0x7e, 0xbd, 0xde, 0xde, 0xf2, 0xde, 0xdc,
	0xeb, 0x5b, 0x4c, 0xa3, 0x3f, 0x1c, 0x2f, 0xed, 0xac, 0xd2, 0x06, 0xd3, 0xf8, 0x7c, 0xbc, 0x18,
	0xab, 0xef, 0xae, 0x20, 0xf2, 0xe4, 0x23, 0x09, 0xb4, 0x37, 0x56, 0x75, 0xcb, 0x5e, 0x63, 0x1a,
	0x77, 0x3c, 0xed, 0x48, 0xc0, 0x6c, 0x7f, 0xa7, 0x12, 0xb0, 0x9d, 0x5b, 0xd5, 0x2d, 0x1e, 0xd6,
	0xee, 0x7f, 0xa5, 0x7b, 0xa9, 0x91, 0x55, 0xdb, 0x07, 0x7c, 0x9c, 0xbf, 0xcf, 0x78, 0x18, 0x9a,
	0xda, 0xba, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xab, 0x05, 0x2f, 0x55, 0xed, 0x00, 0x00, 0x00,
}
