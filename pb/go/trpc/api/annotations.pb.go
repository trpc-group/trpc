// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: trpc/api/annotations.proto

package api

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_trpc_api_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*HttpRule)(nil),
		Field:         50201,
		Name:          "trpc.api.http",
		Tag:           "bytes,50201,opt,name=http",
		Filename:      "trpc/api/annotations.proto",
	},
}

// Extension fields to descriptor.MethodOptions.
var (
	// optional trpc.api.HttpRule http = 50201;
	E_Http = &file_trpc_api_annotations_proto_extTypes[0]
)

var File_trpc_api_annotations_proto protoreflect.FileDescriptor

var file_trpc_api_annotations_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x13, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x48, 0x0a,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x99, 0x88, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x42, 0x4e, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x78, 0x74, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x2e, 0x77, 0x6f,
	0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x72, 0x70, 0x63, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_trpc_api_annotations_proto_goTypes = []interface{}{
	(*descriptor.MethodOptions)(nil), // 0: google.protobuf.MethodOptions
	(*HttpRule)(nil),                 // 1: trpc.api.HttpRule
}
var file_trpc_api_annotations_proto_depIdxs = []int32{
	0, // 0: trpc.api.http:extendee -> google.protobuf.MethodOptions
	1, // 1: trpc.api.http:type_name -> trpc.api.HttpRule
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trpc_api_annotations_proto_init() }
func file_trpc_api_annotations_proto_init() {
	if File_trpc_api_annotations_proto != nil {
		return
	}
	file_trpc_api_http_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trpc_api_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_trpc_api_annotations_proto_goTypes,
		DependencyIndexes: file_trpc_api_annotations_proto_depIdxs,
		ExtensionInfos:    file_trpc_api_annotations_proto_extTypes,
	}.Build()
	File_trpc_api_annotations_proto = out.File
	file_trpc_api_annotations_proto_rawDesc = nil
	file_trpc_api_annotations_proto_goTypes = nil
	file_trpc_api_annotations_proto_depIdxs = nil
}
