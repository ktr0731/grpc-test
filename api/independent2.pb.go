// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: independent2.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FailureCode int32

const (
	FailureCode_FOO FailureCode = 0
	FailureCode_BAR FailureCode = 1
)

// Enum value maps for FailureCode.
var (
	FailureCode_name = map[int32]string{
		0: "FOO",
		1: "BAR",
	}
	FailureCode_value = map[string]int32{
		"FOO": 0,
		"BAR": 1,
	}
)

func (x FailureCode) Enum() *FailureCode {
	p := new(FailureCode)
	*p = x
	return p
}

func (x FailureCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FailureCode) Descriptor() protoreflect.EnumDescriptor {
	return file_independent2_proto_enumTypes[0].Descriptor()
}

func (FailureCode) Type() protoreflect.EnumType {
	return &file_independent2_proto_enumTypes[0]
}

func (x FailureCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FailureCode.Descriptor instead.
func (FailureCode) EnumDescriptor() ([]byte, []int) {
	return file_independent2_proto_rawDescGZIP(), []int{0}
}

var File_independent2_proto protoreflect.FileDescriptor

var file_independent2_proto_rawDesc = []byte{
	0x0a, 0x12, 0x69, 0x6e, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x32, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x2a, 0x1f, 0x0a, 0x0b, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x4f, 0x4f, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x41, 0x52, 0x10, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x74, 0x72, 0x30, 0x37, 0x33, 0x31,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_independent2_proto_rawDescOnce sync.Once
	file_independent2_proto_rawDescData = file_independent2_proto_rawDesc
)

func file_independent2_proto_rawDescGZIP() []byte {
	file_independent2_proto_rawDescOnce.Do(func() {
		file_independent2_proto_rawDescData = protoimpl.X.CompressGZIP(file_independent2_proto_rawDescData)
	})
	return file_independent2_proto_rawDescData
}

var file_independent2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_independent2_proto_goTypes = []interface{}{
	(FailureCode)(0), // 0: api.FailureCode
}
var file_independent2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_independent2_proto_init() }
func file_independent2_proto_init() {
	if File_independent2_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_independent2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_independent2_proto_goTypes,
		DependencyIndexes: file_independent2_proto_depIdxs,
		EnumInfos:         file_independent2_proto_enumTypes,
	}.Build()
	File_independent2_proto = out.File
	file_independent2_proto_rawDesc = nil
	file_independent2_proto_goTypes = nil
	file_independent2_proto_depIdxs = nil
}
