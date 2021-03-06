// Code generated by protoc-gen-go. DO NOT EDIT.
// source: processor_message.proto

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

type CPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Core                 uint32   `protobuf:"varint,3,opt,name=core,proto3" json:"core,omitempty"`
	Thread               uint32   `protobuf:"varint,4,opt,name=thread,proto3" json:"thread,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,5,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,6,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{0}
}

func (m *CPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPU.Unmarshal(m, b)
}
func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
}
func (m *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(m, src)
}
func (m *CPU) XXX_Size() int {
	return xxx_messageInfo_CPU.Size(m)
}
func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CPU) GetCore() uint32 {
	if m != nil {
		return m.Core
	}
	return 0
}

func (m *CPU) GetThread() uint32 {
	if m != nil {
		return m.Thread
	}
	return 0
}

func (m *CPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *CPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

type GPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,3,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,4,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	Memory               *Memory  `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GPU) Reset()         { *m = GPU{} }
func (m *GPU) String() string { return proto.CompactTextString(m) }
func (*GPU) ProtoMessage()    {}
func (*GPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{1}
}

func (m *GPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GPU.Unmarshal(m, b)
}
func (m *GPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GPU.Marshal(b, m, deterministic)
}
func (m *GPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPU.Merge(m, src)
}
func (m *GPU) XXX_Size() int {
	return xxx_messageInfo_GPU.Size(m)
}
func (m *GPU) XXX_DiscardUnknown() {
	xxx_messageInfo_GPU.DiscardUnknown(m)
}

var xxx_messageInfo_GPU proto.InternalMessageInfo

func (m *GPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *GPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *GPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

func (m *GPU) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterType((*CPU)(nil), "laptopstore.pcbook.CPU")
	proto.RegisterType((*GPU)(nil), "laptopstore.pcbook.GPU")
}

func init() {
	proto.RegisterFile("processor_message.proto", fileDescriptor_466578cecc6db379)
}

var fileDescriptor_466578cecc6db379 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xbf, 0x4b, 0xc3, 0x40,
	0x14, 0xe6, 0x9a, 0xf4, 0xc4, 0x13, 0x97, 0x47, 0xb1, 0x47, 0xa7, 0xd0, 0x29, 0x53, 0x86, 0xfa,
	0x1f, 0xe8, 0x90, 0x49, 0x90, 0x80, 0x8b, 0x4b, 0xb9, 0xa4, 0x8f, 0xa6, 0xe8, 0xe5, 0x1d, 0xef,
	0x6e, 0xa8, 0xf9, 0x1f, 0x1c, 0xfd, 0x7f, 0xc5, 0xd7, 0x0c, 0x15, 0x71, 0xe8, 0xf6, 0xfd, 0x82,
	0xef, 0xe3, 0x33, 0xcb, 0xc0, 0xd4, 0x61, 0x8c, 0xc4, 0x5b, 0x8f, 0x31, 0xba, 0x3d, 0x56, 0x81,
	0x29, 0x11, 0xc0, 0xbb, 0x0b, 0x89, 0x42, 0x4c, 0xc4, 0x58, 0x85, 0xae, 0x25, 0x7a, 0x5b, 0x2d,
	0x3c, 0x7a, 0xe2, 0x8f, 0xdf, 0xc9, 0xf5, 0xa7, 0x32, 0xd9, 0xe3, 0xf3, 0x0b, 0x2c, 0xcc, 0xbc,
	0x65, 0x37, 0xec, 0xac, 0x2a, 0x54, 0x79, 0xdd, 0x9c, 0x08, 0x80, 0xc9, 0x07, 0xe7, 0xd1, 0xce,
	0x44, 0x14, 0xfc, 0xa3, 0x75, 0xc4, 0x68, 0xb3, 0x42, 0x95, 0xb7, 0x8d, 0x60, 0xb8, 0x33, 0x3a,
	0xf5, 0x8c, 0x6e, 0x67, 0x73, 0x51, 0x27, 0x06, 0x4b, 0x73, 0xe5, 0x0f, 0xc3, 0x76, 0xdf, 0x8f,
	0x76, 0x5e, 0xa8, 0x52, 0x35, 0xda, 0x1f, 0x86, 0xba, 0x1f, 0xc5, 0x70, 0x47, 0x31, 0xf4, 0x64,
	0xb8, 0x63, 0xdd, 0x8f, 0xeb, 0x2f, 0x65, 0xb2, 0xfa, 0xa2, 0x3d, 0x67, 0x1d, 0xd9, 0x7f, 0x1d,
	0xf9, 0x79, 0x07, 0x6c, 0x8c, 0x3e, 0x7d, 0x21, 0xa3, 0x6e, 0x36, 0xab, 0xea, 0xef, 0x5d, 0xd5,
	0x93, 0x24, 0x9a, 0x29, 0xf9, 0x90, 0xbf, 0xce, 0x42, 0xdb, 0x6a, 0x39, 0xed, 0xfe, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x53, 0xba, 0x67, 0xcd, 0x79, 0x01, 0x00, 0x00,
}
