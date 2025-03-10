// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.3
// source: api/v1/service_account_service.proto

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

// A list of service accounts (free of scoped information)
// Next Tag: 2
type ListServiceAccountResponse struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	SaAndRoles    []*ServiceAccountAndRoles `protobuf:"bytes,1,rep,name=sa_and_roles,json=saAndRoles,proto3" json:"sa_and_roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListServiceAccountResponse) Reset() {
	*x = ListServiceAccountResponse{}
	mi := &file_api_v1_service_account_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListServiceAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceAccountResponse) ProtoMessage() {}

func (x *ListServiceAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_service_account_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceAccountResponse.ProtoReflect.Descriptor instead.
func (*ListServiceAccountResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_service_account_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListServiceAccountResponse) GetSaAndRoles() []*ServiceAccountAndRoles {
	if x != nil {
		return x.SaAndRoles
	}
	return nil
}

// A service account and the roles that reference it
// Next Tag: 5
type ServiceAccountAndRoles struct {
	state                   protoimpl.MessageState      `protogen:"open.v1"`
	ServiceAccount          *storage.ServiceAccount     `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	ClusterRoles            []*storage.K8SRole          `protobuf:"bytes,2,rep,name=cluster_roles,json=clusterRoles,proto3" json:"cluster_roles,omitempty"`
	ScopedRoles             []*ScopedRoles              `protobuf:"bytes,3,rep,name=scoped_roles,json=scopedRoles,proto3" json:"scoped_roles,omitempty"`
	DeploymentRelationships []*SADeploymentRelationship `protobuf:"bytes,4,rep,name=deployment_relationships,json=deploymentRelationships,proto3" json:"deployment_relationships,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *ServiceAccountAndRoles) Reset() {
	*x = ServiceAccountAndRoles{}
	mi := &file_api_v1_service_account_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceAccountAndRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccountAndRoles) ProtoMessage() {}

func (x *ServiceAccountAndRoles) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_service_account_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccountAndRoles.ProtoReflect.Descriptor instead.
func (*ServiceAccountAndRoles) Descriptor() ([]byte, []int) {
	return file_api_v1_service_account_service_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceAccountAndRoles) GetServiceAccount() *storage.ServiceAccount {
	if x != nil {
		return x.ServiceAccount
	}
	return nil
}

func (x *ServiceAccountAndRoles) GetClusterRoles() []*storage.K8SRole {
	if x != nil {
		return x.ClusterRoles
	}
	return nil
}

func (x *ServiceAccountAndRoles) GetScopedRoles() []*ScopedRoles {
	if x != nil {
		return x.ScopedRoles
	}
	return nil
}

func (x *ServiceAccountAndRoles) GetDeploymentRelationships() []*SADeploymentRelationship {
	if x != nil {
		return x.DeploymentRelationships
	}
	return nil
}

// One service account
// Next Tag: 2
type GetServiceAccountResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	SaAndRole     *ServiceAccountAndRoles `protobuf:"bytes,1,opt,name=sa_and_role,json=saAndRole,proto3" json:"sa_and_role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServiceAccountResponse) Reset() {
	*x = GetServiceAccountResponse{}
	mi := &file_api_v1_service_account_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServiceAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceAccountResponse) ProtoMessage() {}

func (x *GetServiceAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_service_account_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceAccountResponse.ProtoReflect.Descriptor instead.
func (*GetServiceAccountResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_service_account_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetServiceAccountResponse) GetSaAndRole() *ServiceAccountAndRoles {
	if x != nil {
		return x.SaAndRole
	}
	return nil
}

// Service accounts can be used by a deployment.
// Next Tag: 3
type SADeploymentRelationship struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of the deployment using the service account
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the deployment.
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SADeploymentRelationship) Reset() {
	*x = SADeploymentRelationship{}
	mi := &file_api_v1_service_account_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SADeploymentRelationship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SADeploymentRelationship) ProtoMessage() {}

func (x *SADeploymentRelationship) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_service_account_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SADeploymentRelationship.ProtoReflect.Descriptor instead.
func (*SADeploymentRelationship) Descriptor() ([]byte, []int) {
	return file_api_v1_service_account_service_proto_rawDescGZIP(), []int{3}
}

func (x *SADeploymentRelationship) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SADeploymentRelationship) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_v1_service_account_service_proto protoreflect.FileDescriptor

var file_api_v1_service_account_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72,
	0x62, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x61, 0x5f, 0x61, 0x6e, 0x64,
	0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x0a, 0x73, 0x61, 0x41, 0x6e, 0x64, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x22, 0x9e, 0x02, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x40, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x35, 0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x64, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52,
	0x0b, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x18,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x41, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x17, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x73, 0x22, 0x57, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x73, 0x61, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x09, 0x73, 0x61, 0x41, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x3e,
	0x0a, 0x18, 0x53, 0x41, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xe1,
	0x01, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a,
	0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x60, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b,
	0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x03, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_service_account_service_proto_rawDescOnce sync.Once
	file_api_v1_service_account_service_proto_rawDescData = file_api_v1_service_account_service_proto_rawDesc
)

func file_api_v1_service_account_service_proto_rawDescGZIP() []byte {
	file_api_v1_service_account_service_proto_rawDescOnce.Do(func() {
		file_api_v1_service_account_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_service_account_service_proto_rawDescData)
	})
	return file_api_v1_service_account_service_proto_rawDescData
}

var file_api_v1_service_account_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_service_account_service_proto_goTypes = []any{
	(*ListServiceAccountResponse)(nil), // 0: v1.ListServiceAccountResponse
	(*ServiceAccountAndRoles)(nil),     // 1: v1.ServiceAccountAndRoles
	(*GetServiceAccountResponse)(nil),  // 2: v1.GetServiceAccountResponse
	(*SADeploymentRelationship)(nil),   // 3: v1.SADeploymentRelationship
	(*storage.ServiceAccount)(nil),     // 4: storage.ServiceAccount
	(*storage.K8SRole)(nil),            // 5: storage.K8sRole
	(*ScopedRoles)(nil),                // 6: v1.ScopedRoles
	(*ResourceByID)(nil),               // 7: v1.ResourceByID
	(*RawQuery)(nil),                   // 8: v1.RawQuery
}
var file_api_v1_service_account_service_proto_depIdxs = []int32{
	1, // 0: v1.ListServiceAccountResponse.sa_and_roles:type_name -> v1.ServiceAccountAndRoles
	4, // 1: v1.ServiceAccountAndRoles.service_account:type_name -> storage.ServiceAccount
	5, // 2: v1.ServiceAccountAndRoles.cluster_roles:type_name -> storage.K8sRole
	6, // 3: v1.ServiceAccountAndRoles.scoped_roles:type_name -> v1.ScopedRoles
	3, // 4: v1.ServiceAccountAndRoles.deployment_relationships:type_name -> v1.SADeploymentRelationship
	1, // 5: v1.GetServiceAccountResponse.sa_and_role:type_name -> v1.ServiceAccountAndRoles
	7, // 6: v1.ServiceAccountService.GetServiceAccount:input_type -> v1.ResourceByID
	8, // 7: v1.ServiceAccountService.ListServiceAccounts:input_type -> v1.RawQuery
	2, // 8: v1.ServiceAccountService.GetServiceAccount:output_type -> v1.GetServiceAccountResponse
	0, // 9: v1.ServiceAccountService.ListServiceAccounts:output_type -> v1.ListServiceAccountResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1_service_account_service_proto_init() }
func file_api_v1_service_account_service_proto_init() {
	if File_api_v1_service_account_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_rbac_service_proto_init()
	file_api_v1_search_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_service_account_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_service_account_service_proto_goTypes,
		DependencyIndexes: file_api_v1_service_account_service_proto_depIdxs,
		MessageInfos:      file_api_v1_service_account_service_proto_msgTypes,
	}.Build()
	File_api_v1_service_account_service_proto = out.File
	file_api_v1_service_account_service_proto_rawDesc = nil
	file_api_v1_service_account_service_proto_goTypes = nil
	file_api_v1_service_account_service_proto_depIdxs = nil
}
