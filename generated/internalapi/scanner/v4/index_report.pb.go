// This contains protobuf types in pair with ClairCore's types. See
// https://github.com/quay/claircore for comments on the fields.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.3
// source: internalapi/scanner/v4/index_report.proto

package v4

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

type IndexReport struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HashId        string                 `protobuf:"bytes,1,opt,name=hash_id,json=hashId,proto3" json:"hash_id,omitempty"`
	State         string                 `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Success       bool                   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Err           string                 `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	Contents      *Contents              `protobuf:"bytes,5,opt,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexReport) Reset() {
	*x = IndexReport{}
	mi := &file_internalapi_scanner_v4_index_report_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexReport) ProtoMessage() {}

func (x *IndexReport) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_index_report_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexReport.ProtoReflect.Descriptor instead.
func (*IndexReport) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_index_report_proto_rawDescGZIP(), []int{0}
}

func (x *IndexReport) GetHashId() string {
	if x != nil {
		return x.HashId
	}
	return ""
}

func (x *IndexReport) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *IndexReport) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *IndexReport) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *IndexReport) GetContents() *Contents {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_internalapi_scanner_v4_index_report_proto protoreflect.FileDescriptor

var file_internalapi_scanner_v4_index_report_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x63, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x1a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a,
	0x0b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x68, 0x61, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68,
	0x61, 0x73, 0x68, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x2f, 0x76, 0x34, 0x3b, 0x76, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_scanner_v4_index_report_proto_rawDescOnce sync.Once
	file_internalapi_scanner_v4_index_report_proto_rawDescData = file_internalapi_scanner_v4_index_report_proto_rawDesc
)

func file_internalapi_scanner_v4_index_report_proto_rawDescGZIP() []byte {
	file_internalapi_scanner_v4_index_report_proto_rawDescOnce.Do(func() {
		file_internalapi_scanner_v4_index_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_scanner_v4_index_report_proto_rawDescData)
	})
	return file_internalapi_scanner_v4_index_report_proto_rawDescData
}

var file_internalapi_scanner_v4_index_report_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internalapi_scanner_v4_index_report_proto_goTypes = []any{
	(*IndexReport)(nil), // 0: scanner.v4.IndexReport
	(*Contents)(nil),    // 1: scanner.v4.Contents
}
var file_internalapi_scanner_v4_index_report_proto_depIdxs = []int32{
	1, // 0: scanner.v4.IndexReport.contents:type_name -> scanner.v4.Contents
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internalapi_scanner_v4_index_report_proto_init() }
func file_internalapi_scanner_v4_index_report_proto_init() {
	if File_internalapi_scanner_v4_index_report_proto != nil {
		return
	}
	file_internalapi_scanner_v4_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_scanner_v4_index_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_scanner_v4_index_report_proto_goTypes,
		DependencyIndexes: file_internalapi_scanner_v4_index_report_proto_depIdxs,
		MessageInfos:      file_internalapi_scanner_v4_index_report_proto_msgTypes,
	}.Build()
	File_internalapi_scanner_v4_index_report_proto = out.File
	file_internalapi_scanner_v4_index_report_proto_rawDesc = nil
	file_internalapi_scanner_v4_index_report_proto_goTypes = nil
	file_internalapi_scanner_v4_index_report_proto_depIdxs = nil
}
