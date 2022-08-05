// Code generated by protoc-gen-go. DO NOT EDIT.
// source: swagger.proto

package trpc

import (
	"fmt"
	"math"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
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

// to gen swagger json
type SwaggerRule struct {
	Title                string          `protobuf:"bytes,50103,opt,name=title,proto3" json:"title,omitempty"`
	Method               string          `protobuf:"bytes,50104,opt,name=method,proto3" json:"method,omitempty"`
	Description          string          `protobuf:"bytes,50105,opt,name=description,proto3" json:"description,omitempty"`
	Params               []*SwaggerParam `protobuf:"bytes,50106,rep,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SwaggerRule) Reset()         { *m = SwaggerRule{} }
func (m *SwaggerRule) String() string { return proto.CompactTextString(m) }
func (*SwaggerRule) ProtoMessage()    {}
func (*SwaggerRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_49635b75e059a131, []int{0}
}

func (m *SwaggerRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwaggerRule.Unmarshal(m, b)
}
func (m *SwaggerRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwaggerRule.Marshal(b, m, deterministic)
}
func (m *SwaggerRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwaggerRule.Merge(m, src)
}
func (m *SwaggerRule) XXX_Size() int {
	return xxx_messageInfo_SwaggerRule.Size(m)
}
func (m *SwaggerRule) XXX_DiscardUnknown() {
	xxx_messageInfo_SwaggerRule.DiscardUnknown(m)
}

var xxx_messageInfo_SwaggerRule proto.InternalMessageInfo

func (m *SwaggerRule) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SwaggerRule) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *SwaggerRule) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SwaggerRule) GetParams() []*SwaggerParam {
	if m != nil {
		return m.Params
	}
	return nil
}

type SwaggerParam struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required             bool     `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	Default              string   `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwaggerParam) Reset()         { *m = SwaggerParam{} }
func (m *SwaggerParam) String() string { return proto.CompactTextString(m) }
func (*SwaggerParam) ProtoMessage()    {}
func (*SwaggerParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_49635b75e059a131, []int{1}
}

func (m *SwaggerParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwaggerParam.Unmarshal(m, b)
}
func (m *SwaggerParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwaggerParam.Marshal(b, m, deterministic)
}
func (m *SwaggerParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwaggerParam.Merge(m, src)
}
func (m *SwaggerParam) XXX_Size() int {
	return xxx_messageInfo_SwaggerParam.Size(m)
}
func (m *SwaggerParam) XXX_DiscardUnknown() {
	xxx_messageInfo_SwaggerParam.DiscardUnknown(m)
}

var xxx_messageInfo_SwaggerParam proto.InternalMessageInfo

func (m *SwaggerParam) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SwaggerParam) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *SwaggerParam) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

var E_Swagger = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MethodOptions)(nil),
	ExtensionType: (*SwaggerRule)(nil),
	Field:         50101,
	Name:          "trpc.swagger",
	Tag:           "bytes,50101,opt,name=swagger",
	Filename:      "swagger.proto",
}

func init() {
	proto.RegisterType((*SwaggerRule)(nil), "trpc.SwaggerRule")
	proto.RegisterType((*SwaggerParam)(nil), "trpc.SwaggerParam")
	proto.RegisterExtension(E_Swagger)
}

func init() { proto.RegisterFile("swagger.proto", fileDescriptor_49635b75e059a131) }

var fileDescriptor_49635b75e059a131 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x50, 0x3f, 0x4f, 0xfb, 0x30,
	0x14, 0x54, 0x7e, 0xed, 0xaf, 0x2d, 0x2f, 0x30, 0xe0, 0x01, 0x59, 0x15, 0xa0, 0xa8, 0x53, 0x06,
	0x70, 0xa4, 0xb2, 0x21, 0xb1, 0xb0, 0x57, 0xa0, 0xb0, 0x20, 0x36, 0xd7, 0x79, 0x35, 0x91, 0x92,
	0xbc, 0xe0, 0x38, 0x82, 0x6f, 0xd0, 0x99, 0x0f, 0x83, 0xf8, 0xf3, 0xe9, 0x50, 0x1c, 0xa7, 0x2a,
	0x93, 0x7d, 0xef, 0xce, 0xe7, 0x7b, 0x07, 0x47, 0xcd, 0xab, 0xd4, 0x1a, 0x8d, 0xa8, 0x0d, 0x59,
	0x62, 0x63, 0x6b, 0x6a, 0x35, 0x8f, 0x34, 0x91, 0x2e, 0x30, 0x71, 0xb3, 0x75, 0xbb, 0x49, 0x32,
	0x6c, 0x94, 0xc9, 0x6b, 0x4b, 0x5e, 0xb7, 0x78, 0x0f, 0x20, 0x7c, 0xe8, 0x5f, 0xa6, 0x6d, 0x81,
	0xec, 0x04, 0xfe, 0xdb, 0xdc, 0x16, 0xc8, 0x3f, 0xb7, 0xa3, 0x28, 0x88, 0x0f, 0xd2, 0x1e, 0x32,
	0x0e, 0x93, 0x12, 0xed, 0x33, 0x65, 0xfc, 0xcb, 0x13, 0x1e, 0xb3, 0x05, 0x84, 0x83, 0x6b, 0x4e,
	0x15, 0xff, 0xf6, 0xf4, 0xfe, 0x90, 0x5d, 0xc0, 0xa4, 0x96, 0x46, 0x96, 0x0d, 0xff, 0xd9, 0x8e,
	0xa2, 0x51, 0x1c, 0x2e, 0x99, 0xe8, 0x02, 0x0a, 0xff, 0xf5, 0x7d, 0x47, 0xa6, 0x5e, 0xb3, 0x78,
	0x84, 0xc3, 0xfd, 0x39, 0x63, 0x30, 0xae, 0x64, 0x89, 0x3c, 0x70, 0xc6, 0xee, 0xce, 0xe6, 0x30,
	0x33, 0xf8, 0xd2, 0xe6, 0x06, 0x33, 0xfe, 0x2f, 0x0a, 0xe2, 0x59, 0xba, 0xc3, 0x8c, 0xc3, 0x34,
	0xc3, 0x8d, 0x6c, 0x0b, 0xcb, 0xfb, 0x2c, 0x03, 0xbc, 0x5e, 0xc1, 0xd4, 0xd7, 0xc4, 0xce, 0x45,
	0xdf, 0x8d, 0x18, 0xba, 0x11, 0x2b, 0xb7, 0xcf, 0x9d, 0x4b, 0xdc, 0xf0, 0x0f, 0xb7, 0x47, 0xb8,
	0x3c, 0xfe, 0x13, 0xb4, 0xeb, 0x28, 0x1d, 0x3c, 0x6e, 0x6f, 0xe0, 0x4c, 0x51, 0x29, 0x2c, 0x56,
	0x0a, 0x2b, 0xdb, 0x0b, 0x77, 0x76, 0xf8, 0x66, 0x9f, 0x4e, 0x75, 0x6e, 0x85, 0xa2, 0x0c, 0x05,
	0x49, 0xa1, 0xa8, 0x4c, 0x3a, 0xc9, 0xa5, 0x26, 0x77, 0xae, 0x27, 0x4e, 0x7b, 0xf5, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0xf7, 0x02, 0x96, 0xbb, 0x01, 0x00, 0x00,
}
