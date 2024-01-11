// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/errcode.proto

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

type ErrCode int32

const (
	ErrCode_Success ErrCode = 0
	ErrCode_UnKnown ErrCode = 1
)

var ErrCode_name = map[int32]string{
	0: "Success",
	1: "UnKnown",
}

var ErrCode_value = map[string]int32{
	"Success": 0,
	"UnKnown": 1,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8438add6648f46df, []int{0}
}

func init() {
	proto.RegisterEnum("ErrCode", ErrCode_name, ErrCode_value)
}

func init() { proto.RegisterFile("common/errcode.proto", fileDescriptor_8438add6648f46df) }

var fileDescriptor_8438add6648f46df = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x4f, 0x2d, 0x2a, 0x4a, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0xd7, 0x52, 0xe6, 0x62, 0x77, 0x2d, 0x2a, 0x72, 0xce, 0x4f, 0x49, 0x15, 0xe2, 0xe6, 0x62, 0x0f,
	0x2e, 0x4d, 0x4e, 0x4e, 0x2d, 0x2e, 0x16, 0x60, 0x00, 0x71, 0x42, 0xf3, 0xbc, 0xf3, 0xf2, 0xcb,
	0xf3, 0x04, 0x18, 0x9d, 0xb8, 0xa3, 0x58, 0xf4, 0xac, 0x0b, 0x92, 0x56, 0x31, 0x31, 0x05, 0x24,
	0x25, 0xb1, 0x81, 0x35, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xad, 0x57, 0x47, 0x50,
	0x00, 0x00, 0x00,
}
