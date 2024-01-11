// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

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

type Vector3 struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector3) Reset()         { *m = Vector3{} }
func (m *Vector3) String() string { return proto.CompactTextString(m) }
func (*Vector3) ProtoMessage()    {}
func (*Vector3) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *Vector3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector3.Unmarshal(m, b)
}
func (m *Vector3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector3.Marshal(b, m, deterministic)
}
func (m *Vector3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector3.Merge(m, src)
}
func (m *Vector3) XXX_Size() int {
	return xxx_messageInfo_Vector3.Size(m)
}
func (m *Vector3) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector3.DiscardUnknown(m)
}

var xxx_messageInfo_Vector3 proto.InternalMessageInfo

func (m *Vector3) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector3) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vector3) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

type Entity struct {
	RoleId               uint32   `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	RoleMask             int32    `protobuf:"varint,2,opt,name=roleMask,proto3" json:"roleMask,omitempty"`
	Position             *Vector3 `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetRoleId() uint32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *Entity) GetRoleMask() int32 {
	if m != nil {
		return m.RoleMask
	}
	return 0
}

func (m *Entity) GetPosition() *Vector3 {
	if m != nil {
		return m.Position
	}
	return nil
}

func init() {
	proto.RegisterType((*Vector3)(nil), "Vector3")
	proto.RegisterType((*Entity)(nil), "Entity")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x4a, 0xc6, 0x5c, 0xec, 0x61,
	0xa9, 0xc9, 0x25, 0xf9, 0x45, 0xc6, 0x42, 0x3c, 0x5c, 0x8c, 0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x4c, 0x41, 0x8c, 0x15, 0x20, 0x5e, 0xa5, 0x04, 0x13, 0x84, 0x57, 0x09, 0xe2, 0x55, 0x49, 0x30,
	0x43, 0x78, 0x55, 0x4a, 0x49, 0x5c, 0x6c, 0xae, 0x79, 0x25, 0x99, 0x25, 0x95, 0x42, 0x62, 0x5c,
	0x6c, 0x45, 0xf9, 0x39, 0xa9, 0x9e, 0x29, 0x60, 0x8d, 0xbc, 0x41, 0x50, 0x9e, 0x90, 0x14, 0x17,
	0x07, 0x88, 0xe5, 0x9b, 0x58, 0x9c, 0x0d, 0x36, 0x84, 0x35, 0x08, 0xce, 0x17, 0x52, 0xe1, 0xe2,
	0x28, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0xcc, 0xcf, 0x03, 0x1b, 0xc9, 0x6d, 0xc4, 0xa1, 0x07, 0x75,
	0x43, 0x10, 0x5c, 0xc6, 0x89, 0x3b, 0x8a, 0x45, 0xcf, 0xba, 0x20, 0x69, 0x15, 0x13, 0x53, 0x40,
	0x52, 0x12, 0x1b, 0xd8, 0xb1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x04, 0x24, 0xf1,
	0xc3, 0x00, 0x00, 0x00,
}
