// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sometype.proto

package sometype

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	example2 "go_gen/proto3cloneV2/example2"
	example3 "go_gen/proto3cloneV2/example3"
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

type ApplesServerin struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Student              []int32  `protobuf:"varint,3,rep,packed,name=student,proto3" json:"student,omitempty"`
	Teacher              []string `protobuf:"bytes,4,rep,name=teacher,proto3" json:"teacher,omitempty"`
	Bb                   *Binary  `protobuf:"bytes,5,opt,name=bb,proto3" json:"bb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplesServerin) Reset()         { *m = ApplesServerin{} }
func (m *ApplesServerin) String() string { return proto.CompactTextString(m) }
func (*ApplesServerin) ProtoMessage()    {}
func (*ApplesServerin) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddf4853b7126b944, []int{0}
}

func (m *ApplesServerin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplesServerin.Unmarshal(m, b)
}
func (m *ApplesServerin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplesServerin.Marshal(b, m, deterministic)
}
func (m *ApplesServerin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplesServerin.Merge(m, src)
}
func (m *ApplesServerin) XXX_Size() int {
	return xxx_messageInfo_ApplesServerin.Size(m)
}
func (m *ApplesServerin) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplesServerin.DiscardUnknown(m)
}

var xxx_messageInfo_ApplesServerin proto.InternalMessageInfo

func (m *ApplesServerin) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ApplesServerin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplesServerin) GetStudent() []int32 {
	if m != nil {
		return m.Student
	}
	return nil
}

func (m *ApplesServerin) GetTeacher() []string {
	if m != nil {
		return m.Teacher
	}
	return nil
}

func (m *ApplesServerin) GetBb() *Binary {
	if m != nil {
		return m.Bb
	}
	return nil
}

type Binary struct {
	Id                   []string          `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Hero                 *example2.Hero    `protobuf:"bytes,2,opt,name=hero,proto3" json:"hero,omitempty"`
	Tt                   *example3.Teacher `protobuf:"bytes,3,opt,name=tt,proto3" json:"tt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Binary) Reset()         { *m = Binary{} }
func (m *Binary) String() string { return proto.CompactTextString(m) }
func (*Binary) ProtoMessage()    {}
func (*Binary) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddf4853b7126b944, []int{1}
}

func (m *Binary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Binary.Unmarshal(m, b)
}
func (m *Binary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Binary.Marshal(b, m, deterministic)
}
func (m *Binary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Binary.Merge(m, src)
}
func (m *Binary) XXX_Size() int {
	return xxx_messageInfo_Binary.Size(m)
}
func (m *Binary) XXX_DiscardUnknown() {
	xxx_messageInfo_Binary.DiscardUnknown(m)
}

var xxx_messageInfo_Binary proto.InternalMessageInfo

func (m *Binary) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Binary) GetHero() *example2.Hero {
	if m != nil {
		return m.Hero
	}
	return nil
}

func (m *Binary) GetTt() *example3.Teacher {
	if m != nil {
		return m.Tt
	}
	return nil
}

func init() {
	proto.RegisterType((*ApplesServerin)(nil), "sometype.applesServerin")
	proto.RegisterType((*Binary)(nil), "sometype.Binary")
}

func init() {
	proto.RegisterFile("sometype.proto", fileDescriptor_ddf4853b7126b944)
}

var fileDescriptor_ddf4853b7126b944 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xb1, 0x6e, 0xc3, 0x20,
	0x14, 0x45, 0x65, 0x6c, 0xa7, 0x0d, 0x91, 0xac, 0x96, 0x09, 0x75, 0xa2, 0x9e, 0x3c, 0xd9, 0x95,
	0x33, 0x76, 0xcb, 0xd4, 0x99, 0x56, 0x1d, 0xba, 0x44, 0x10, 0x3f, 0x25, 0x96, 0x6c, 0x40, 0x98,
	0x56, 0xcd, 0x2f, 0xf4, 0xab, 0x2b, 0x4c, 0x60, 0x7b, 0x47, 0xf7, 0x0a, 0xce, 0xc5, 0xd5, 0xa2,
	0x67, 0x70, 0x57, 0x03, 0xad, 0xb1, 0xda, 0x69, 0x72, 0x1f, 0xf9, 0xa9, 0x82, 0x5f, 0x31, 0x9b,
	0x09, 0xfa, 0x90, 0x24, 0xde, 0x07, 0xae, 0xff, 0x32, 0x5c, 0x09, 0x63, 0x26, 0x58, 0xde, 0xc1,
	0xfe, 0x80, 0x1d, 0x15, 0xa9, 0x30, 0x1a, 0x07, 0x9a, 0xb1, 0xac, 0x29, 0x39, 0x1a, 0x07, 0x42,
	0x70, 0xa1, 0xc4, 0x0c, 0x14, 0xb1, 0xac, 0xd9, 0xf2, 0xf5, 0x26, 0x14, 0xdf, 0x2d, 0xee, 0x7b,
	0x00, 0xe5, 0x68, 0xce, 0xf2, 0xa6, 0xe4, 0x11, 0x7d, 0xe2, 0x40, 0x9c, 0x2e, 0x60, 0x69, 0xc1,
	0xf2, 0x66, 0xcb, 0x23, 0x12, 0x86, 0x91, 0x94, 0xb4, 0x64, 0x59, 0xb3, 0xeb, 0x1f, 0xda, 0x64,
	0x7c, 0x18, 0x95, 0xb0, 0x57, 0x8e, 0xa4, 0xac, 0x8f, 0x78, 0x13, 0x28, 0x39, 0xf8, 0x07, 0xbc,
	0x43, 0x8d, 0x8b, 0x0b, 0x58, 0xbd, 0x3a, 0xec, 0xfa, 0xaa, 0x4d, 0xab, 0xde, 0xc0, 0x6a, 0xbe,
	0x66, 0xe4, 0x19, 0x23, 0xe7, 0x75, 0x7c, 0xe3, 0xb1, 0x4d, 0x3b, 0x3f, 0xc2, 0xf7, 0x1c, 0x39,
	0x77, 0x78, 0xf9, 0x6a, 0xcf, 0xfa, 0x78, 0x06, 0xd5, 0xad, 0xeb, 0xf7, 0xa7, 0x49, 0x2b, 0xf8,
	0xec, 0xbb, 0x5b, 0xb9, 0x8b, 0x52, 0xaf, 0xf1, 0x90, 0x9b, 0x50, 0xfc, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0xb2, 0x9f, 0x88, 0x17, 0x62, 0x01, 0x00, 0x00,
}