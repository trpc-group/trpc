// Tencent is pleased to support the open source community by making tRPC available.
// Copyright (C) 2023 Tencent. All rights reserved.
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: trpc/swagger/swagger.proto

package swagger

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// to gen swagger json
type SwaggerRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string          `protobuf:"bytes,50103,opt,name=title,proto3" json:"title,omitempty"`
	Method      string          `protobuf:"bytes,50104,opt,name=method,proto3" json:"method,omitempty"`
	Description string          `protobuf:"bytes,50105,opt,name=description,proto3" json:"description,omitempty"`
	Params      []*SwaggerParam `protobuf:"bytes,50106,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *SwaggerRule) Reset() {
	*x = SwaggerRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trpc_swagger_swagger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwaggerRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwaggerRule) ProtoMessage() {}

func (x *SwaggerRule) ProtoReflect() protoreflect.Message {
	mi := &file_trpc_swagger_swagger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwaggerRule.ProtoReflect.Descriptor instead.
func (*SwaggerRule) Descriptor() ([]byte, []int) {
	return file_trpc_swagger_swagger_proto_rawDescGZIP(), []int{0}
}

func (x *SwaggerRule) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SwaggerRule) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *SwaggerRule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SwaggerRule) GetParams() []*SwaggerParam {
	if x != nil {
		return x.Params
	}
	return nil
}

type SwaggerParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required bool   `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	Default  string `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *SwaggerParam) Reset() {
	*x = SwaggerParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trpc_swagger_swagger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwaggerParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwaggerParam) ProtoMessage() {}

func (x *SwaggerParam) ProtoReflect() protoreflect.Message {
	mi := &file_trpc_swagger_swagger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwaggerParam.ProtoReflect.Descriptor instead.
func (*SwaggerParam) Descriptor() ([]byte, []int) {
	return file_trpc_swagger_swagger_proto_rawDescGZIP(), []int{1}
}

func (x *SwaggerParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SwaggerParam) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *SwaggerParam) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

var file_trpc_swagger_swagger_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*SwaggerRule)(nil),
		Field:         50101,
		Name:          "trpc.swagger",
		Tag:           "bytes,50101,opt,name=swagger",
		Filename:      "trpc/swagger/swagger.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional trpc.SwaggerRule swagger = 50101;
	E_Swagger = &file_trpc_swagger_swagger_proto_extTypes[0]
)

var File_trpc_swagger_swagger_proto protoreflect.FileDescriptor

var file_trpc_swagger_swagger_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x73,
	0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x72,
	0x70, 0x63, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0xb7, 0x87,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0xb8, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x22, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xb9, 0x87, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0xba, 0x87, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x58, 0x0a, 0x0c, 0x53, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x3a, 0x4d, 0x0a, 0x07, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb5, 0x87,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x42, 0x51, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74,
	0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65,
	0x78, 0x74, 0x5a, 0x30, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x74,
	0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trpc_swagger_swagger_proto_rawDescOnce sync.Once
	file_trpc_swagger_swagger_proto_rawDescData = file_trpc_swagger_swagger_proto_rawDesc
)

func file_trpc_swagger_swagger_proto_rawDescGZIP() []byte {
	file_trpc_swagger_swagger_proto_rawDescOnce.Do(func() {
		file_trpc_swagger_swagger_proto_rawDescData = protoimpl.X.CompressGZIP(file_trpc_swagger_swagger_proto_rawDescData)
	})
	return file_trpc_swagger_swagger_proto_rawDescData
}

var file_trpc_swagger_swagger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_trpc_swagger_swagger_proto_goTypes = []interface{}{
	(*SwaggerRule)(nil),                // 0: trpc.SwaggerRule
	(*SwaggerParam)(nil),               // 1: trpc.SwaggerParam
	(*descriptorpb.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
}
var file_trpc_swagger_swagger_proto_depIdxs = []int32{
	1, // 0: trpc.SwaggerRule.params:type_name -> trpc.SwaggerParam
	2, // 1: trpc.swagger:extendee -> google.protobuf.MethodOptions
	0, // 2: trpc.swagger:type_name -> trpc.SwaggerRule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trpc_swagger_swagger_proto_init() }
func file_trpc_swagger_swagger_proto_init() {
	if File_trpc_swagger_swagger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trpc_swagger_swagger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwaggerRule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_trpc_swagger_swagger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwaggerParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trpc_swagger_swagger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_trpc_swagger_swagger_proto_goTypes,
		DependencyIndexes: file_trpc_swagger_swagger_proto_depIdxs,
		MessageInfos:      file_trpc_swagger_swagger_proto_msgTypes,
		ExtensionInfos:    file_trpc_swagger_swagger_proto_extTypes,
	}.Build()
	File_trpc_swagger_swagger_proto = out.File
	file_trpc_swagger_swagger_proto_rawDesc = nil
	file_trpc_swagger_swagger_proto_goTypes = nil
	file_trpc_swagger_swagger_proto_depIdxs = nil
}
