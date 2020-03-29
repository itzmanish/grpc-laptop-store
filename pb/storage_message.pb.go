// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage_message.proto

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

type Storage_DRIVER int32

const (
	Storage_DEFAULT Storage_DRIVER = 0
	Storage_SSD     Storage_DRIVER = 1
	Storage_HDD     Storage_DRIVER = 2
)

var Storage_DRIVER_name = map[int32]string{
	0: "DEFAULT",
	1: "SSD",
	2: "HDD",
}

var Storage_DRIVER_value = map[string]int32{
	"DEFAULT": 0,
	"SSD":     1,
	"HDD":     2,
}

func (x Storage_DRIVER) String() string {
	return proto.EnumName(Storage_DRIVER_name, int32(x))
}

func (Storage_DRIVER) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0, 0}
}

type Storage struct {
	Driver               Storage_DRIVER `protobuf:"varint,1,opt,name=driver,proto3,enum=laptopstore.pcbook.Storage_DRIVER" json:"driver,omitempty"`
	Memory               *Memory        `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetDriver() Storage_DRIVER {
	if m != nil {
		return m.Driver
	}
	return Storage_DEFAULT
}

func (m *Storage) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterEnum("laptopstore.pcbook.Storage_DRIVER", Storage_DRIVER_name, Storage_DRIVER_value)
	proto.RegisterType((*Storage)(nil), "laptopstore.pcbook.Storage")
}

func init() { proto.RegisterFile("storage_message.proto", fileDescriptor_170f09d838bd8a04) }

var fileDescriptor_170f09d838bd8a04 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0x8d, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0xca, 0x49, 0x2c, 0x28, 0xc9, 0x2f, 0x00, 0x49, 0xa6, 0xea, 0x15, 0x24, 0x27,
	0xe5, 0xe7, 0x67, 0x4b, 0x89, 0xe4, 0xa6, 0xe6, 0xe6, 0x17, 0x55, 0xa2, 0xaa, 0x54, 0x5a, 0xc4,
	0xc8, 0xc5, 0x1e, 0x0c, 0x31, 0x43, 0xc8, 0x8a, 0x8b, 0x2d, 0xa5, 0x28, 0xb3, 0x2c, 0xb5, 0x48,
	0x82, 0x51, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x49, 0x0f, 0xd3, 0x18, 0x3d, 0xa8, 0x62, 0x3d, 0x97,
	0x20, 0xcf, 0x30, 0xd7, 0xa0, 0x20, 0xa8, 0x0e, 0x21, 0x23, 0x2e, 0x36, 0x88, 0xf9, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x52, 0xd8, 0xf4, 0xfa, 0x82, 0x55, 0x04, 0x41, 0x55, 0x2a, 0xa9,
	0x73, 0xb1, 0x41, 0x4c, 0x11, 0xe2, 0xe6, 0x62, 0x77, 0x71, 0x75, 0x73, 0x0c, 0xf5, 0x09, 0x11,
	0x60, 0x10, 0x62, 0xe7, 0x62, 0x0e, 0x0e, 0x76, 0x11, 0x60, 0x04, 0x31, 0x3c, 0x5c, 0x5c, 0x04,
	0x98, 0x9c, 0x58, 0xa2, 0x98, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0x2e, 0x36, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xc2, 0x72, 0xfb, 0xf7, 0xf4, 0x00, 0x00, 0x00,
}