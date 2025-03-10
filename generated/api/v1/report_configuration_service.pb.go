// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.3
// source: api/v1/report_configuration_service.proto

package v1

import (
	storage "github.com/stackrox/rox/generated/storage"
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

type GetReportConfigurationsResponse struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	ReportConfigs []*storage.ReportConfiguration `protobuf:"bytes,1,rep,name=report_configs,json=reportConfigs,proto3" json:"report_configs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReportConfigurationsResponse) Reset() {
	*x = GetReportConfigurationsResponse{}
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReportConfigurationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportConfigurationsResponse) ProtoMessage() {}

func (x *GetReportConfigurationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportConfigurationsResponse.ProtoReflect.Descriptor instead.
func (*GetReportConfigurationsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_report_configuration_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetReportConfigurationsResponse) GetReportConfigs() []*storage.ReportConfiguration {
	if x != nil {
		return x.ReportConfigs
	}
	return nil
}

type GetReportConfigurationResponse struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	ReportConfig  *storage.ReportConfiguration `protobuf:"bytes,1,opt,name=report_config,json=reportConfig,proto3" json:"report_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReportConfigurationResponse) Reset() {
	*x = GetReportConfigurationResponse{}
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReportConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportConfigurationResponse) ProtoMessage() {}

func (x *GetReportConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportConfigurationResponse.ProtoReflect.Descriptor instead.
func (*GetReportConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_report_configuration_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetReportConfigurationResponse) GetReportConfig() *storage.ReportConfiguration {
	if x != nil {
		return x.ReportConfig
	}
	return nil
}

type PostReportConfigurationResponse struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	ReportConfig  *storage.ReportConfiguration `protobuf:"bytes,1,opt,name=report_config,json=reportConfig,proto3" json:"report_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostReportConfigurationResponse) Reset() {
	*x = PostReportConfigurationResponse{}
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostReportConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostReportConfigurationResponse) ProtoMessage() {}

func (x *PostReportConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostReportConfigurationResponse.ProtoReflect.Descriptor instead.
func (*PostReportConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_report_configuration_service_proto_rawDescGZIP(), []int{2}
}

func (x *PostReportConfigurationResponse) GetReportConfig() *storage.ReportConfiguration {
	if x != nil {
		return x.ReportConfig
	}
	return nil
}

type PostReportConfigurationRequest struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	ReportConfig  *storage.ReportConfiguration `protobuf:"bytes,1,opt,name=report_config,json=reportConfig,proto3" json:"report_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostReportConfigurationRequest) Reset() {
	*x = PostReportConfigurationRequest{}
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostReportConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostReportConfigurationRequest) ProtoMessage() {}

func (x *PostReportConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostReportConfigurationRequest.ProtoReflect.Descriptor instead.
func (*PostReportConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_report_configuration_service_proto_rawDescGZIP(), []int{3}
}

func (x *PostReportConfigurationRequest) GetReportConfig() *storage.ReportConfiguration {
	if x != nil {
		return x.ReportConfig
	}
	return nil
}

type UpdateReportConfigurationRequest struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Id            string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReportConfig  *storage.ReportConfiguration `protobuf:"bytes,2,opt,name=report_config,json=reportConfig,proto3" json:"report_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateReportConfigurationRequest) Reset() {
	*x = UpdateReportConfigurationRequest{}
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReportConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReportConfigurationRequest) ProtoMessage() {}

func (x *UpdateReportConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReportConfigurationRequest.ProtoReflect.Descriptor instead.
func (*UpdateReportConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_report_configuration_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateReportConfigurationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateReportConfigurationRequest) GetReportConfig() *storage.ReportConfiguration {
	if x != nil {
		return x.ReportConfig
	}
	return nil
}

type CountReportConfigurationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int32                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CountReportConfigurationsResponse) Reset() {
	*x = CountReportConfigurationsResponse{}
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountReportConfigurationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountReportConfigurationsResponse) ProtoMessage() {}

func (x *CountReportConfigurationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_report_configuration_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountReportConfigurationsResponse.ProtoReflect.Descriptor instead.
func (*CountReportConfigurationsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_report_configuration_service_proto_rawDescGZIP(), []int{5}
}

func (x *CountReportConfigurationsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_api_v1_report_configuration_service_proto protoreflect.FileDescriptor

var file_api_v1_report_configuration_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22,
	0x63, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x64, 0x0a, 0x1f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x63, 0x0a, 0x1e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0d,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x75, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x39, 0x0a, 0x21, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0xe6, 0x05, 0x0a, 0x1a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x76, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x22, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x88, 0x01, 0x0a, 0x17, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x77, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x1a, 0x1e, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x60, 0x0a,
	0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a,
	0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x79, 0x0a, 0x19, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x25, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x58, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_report_configuration_service_proto_rawDescOnce sync.Once
	file_api_v1_report_configuration_service_proto_rawDescData = file_api_v1_report_configuration_service_proto_rawDesc
)

func file_api_v1_report_configuration_service_proto_rawDescGZIP() []byte {
	file_api_v1_report_configuration_service_proto_rawDescOnce.Do(func() {
		file_api_v1_report_configuration_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_report_configuration_service_proto_rawDescData)
	})
	return file_api_v1_report_configuration_service_proto_rawDescData
}

var file_api_v1_report_configuration_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_v1_report_configuration_service_proto_goTypes = []any{
	(*GetReportConfigurationsResponse)(nil),   // 0: v1.GetReportConfigurationsResponse
	(*GetReportConfigurationResponse)(nil),    // 1: v1.GetReportConfigurationResponse
	(*PostReportConfigurationResponse)(nil),   // 2: v1.PostReportConfigurationResponse
	(*PostReportConfigurationRequest)(nil),    // 3: v1.PostReportConfigurationRequest
	(*UpdateReportConfigurationRequest)(nil),  // 4: v1.UpdateReportConfigurationRequest
	(*CountReportConfigurationsResponse)(nil), // 5: v1.CountReportConfigurationsResponse
	(*storage.ReportConfiguration)(nil),       // 6: storage.ReportConfiguration
	(*RawQuery)(nil),                          // 7: v1.RawQuery
	(*ResourceByID)(nil),                      // 8: v1.ResourceByID
	(*Empty)(nil),                             // 9: v1.Empty
}
var file_api_v1_report_configuration_service_proto_depIdxs = []int32{
	6,  // 0: v1.GetReportConfigurationsResponse.report_configs:type_name -> storage.ReportConfiguration
	6,  // 1: v1.GetReportConfigurationResponse.report_config:type_name -> storage.ReportConfiguration
	6,  // 2: v1.PostReportConfigurationResponse.report_config:type_name -> storage.ReportConfiguration
	6,  // 3: v1.PostReportConfigurationRequest.report_config:type_name -> storage.ReportConfiguration
	6,  // 4: v1.UpdateReportConfigurationRequest.report_config:type_name -> storage.ReportConfiguration
	7,  // 5: v1.ReportConfigurationService.GetReportConfigurations:input_type -> v1.RawQuery
	8,  // 6: v1.ReportConfigurationService.GetReportConfiguration:input_type -> v1.ResourceByID
	3,  // 7: v1.ReportConfigurationService.PostReportConfiguration:input_type -> v1.PostReportConfigurationRequest
	4,  // 8: v1.ReportConfigurationService.UpdateReportConfiguration:input_type -> v1.UpdateReportConfigurationRequest
	8,  // 9: v1.ReportConfigurationService.DeleteReportConfiguration:input_type -> v1.ResourceByID
	7,  // 10: v1.ReportConfigurationService.CountReportConfigurations:input_type -> v1.RawQuery
	0,  // 11: v1.ReportConfigurationService.GetReportConfigurations:output_type -> v1.GetReportConfigurationsResponse
	1,  // 12: v1.ReportConfigurationService.GetReportConfiguration:output_type -> v1.GetReportConfigurationResponse
	2,  // 13: v1.ReportConfigurationService.PostReportConfiguration:output_type -> v1.PostReportConfigurationResponse
	9,  // 14: v1.ReportConfigurationService.UpdateReportConfiguration:output_type -> v1.Empty
	9,  // 15: v1.ReportConfigurationService.DeleteReportConfiguration:output_type -> v1.Empty
	5,  // 16: v1.ReportConfigurationService.CountReportConfigurations:output_type -> v1.CountReportConfigurationsResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_report_configuration_service_proto_init() }
func file_api_v1_report_configuration_service_proto_init() {
	if File_api_v1_report_configuration_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	file_api_v1_search_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_report_configuration_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_report_configuration_service_proto_goTypes,
		DependencyIndexes: file_api_v1_report_configuration_service_proto_depIdxs,
		MessageInfos:      file_api_v1_report_configuration_service_proto_msgTypes,
	}.Build()
	File_api_v1_report_configuration_service_proto = out.File
	file_api_v1_report_configuration_service_proto_rawDesc = nil
	file_api_v1_report_configuration_service_proto_goTypes = nil
	file_api_v1_report_configuration_service_proto_depIdxs = nil
}
