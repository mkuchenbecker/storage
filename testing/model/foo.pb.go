// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo.proto

package testing_model

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

type Foo struct {
	Bar                  string   `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}

func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

func init() {
	proto.RegisterType((*Foo)(nil), "storage.client.Foo")
}

func init() {
	proto.RegisterFile("foo.proto", fileDescriptor_7ce1e2eec643ca48)
}

var fileDescriptor_7ce1e2eec643ca48 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x4b,
	0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x51, 0x12, 0xe7, 0x62, 0x76, 0xcb, 0xcf, 0x17, 0x12, 0xe0, 0x62,
	0x4e, 0x4a, 0x2c, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x9d, 0xf8, 0xa3, 0x78,
	0x4b, 0x52, 0x8b, 0x4b, 0x32, 0xf3, 0xd2, 0xf5, 0x72, 0xf3, 0x53, 0x52, 0x73, 0x92, 0xd8, 0xc0,
	0x06, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xd5, 0xab, 0xe3, 0x4d, 0x00, 0x00, 0x00,
}
