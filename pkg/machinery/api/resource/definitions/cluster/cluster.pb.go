// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: resource/definitions/cluster/cluster.proto

package cluster

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	common "github.com/siderolabs/talos/pkg/machinery/api/common"
	enums "github.com/siderolabs/talos/pkg/machinery/api/resource/definitions/enums"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AffiliateSpec describes Affiliate state.
type AffiliateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId          string                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Addresses       []*common.NetIP        `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Hostname        string                 `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Nodename        string                 `protobuf:"bytes,4,opt,name=nodename,proto3" json:"nodename,omitempty"`
	OperatingSystem string                 `protobuf:"bytes,5,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
	MachineType     enums.MachineType      `protobuf:"varint,6,opt,name=machine_type,json=machineType,proto3,enum=talos.resource.definitions.enums.MachineType" json:"machine_type,omitempty"`
	KubeSpan        *KubeSpanAffiliateSpec `protobuf:"bytes,7,opt,name=kube_span,json=kubeSpan,proto3" json:"kube_span,omitempty"`
	ControlPlane    *ControlPlane          `protobuf:"bytes,8,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
}

func (x *AffiliateSpec) Reset() {
	*x = AffiliateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AffiliateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AffiliateSpec) ProtoMessage() {}

func (x *AffiliateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AffiliateSpec.ProtoReflect.Descriptor instead.
func (*AffiliateSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_cluster_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *AffiliateSpec) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *AffiliateSpec) GetAddresses() []*common.NetIP {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *AffiliateSpec) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *AffiliateSpec) GetNodename() string {
	if x != nil {
		return x.Nodename
	}
	return ""
}

func (x *AffiliateSpec) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *AffiliateSpec) GetMachineType() enums.MachineType {
	if x != nil {
		return x.MachineType
	}
	return enums.MachineType(0)
}

func (x *AffiliateSpec) GetKubeSpan() *KubeSpanAffiliateSpec {
	if x != nil {
		return x.KubeSpan
	}
	return nil
}

func (x *AffiliateSpec) GetControlPlane() *ControlPlane {
	if x != nil {
		return x.ControlPlane
	}
	return nil
}

// ConfigSpec describes KubeSpan configuration.
type ConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscoveryEnabled          bool   `protobuf:"varint,1,opt,name=discovery_enabled,json=discoveryEnabled,proto3" json:"discovery_enabled,omitempty"`
	RegistryKubernetesEnabled bool   `protobuf:"varint,2,opt,name=registry_kubernetes_enabled,json=registryKubernetesEnabled,proto3" json:"registry_kubernetes_enabled,omitempty"`
	RegistryServiceEnabled    bool   `protobuf:"varint,3,opt,name=registry_service_enabled,json=registryServiceEnabled,proto3" json:"registry_service_enabled,omitempty"`
	ServiceEndpoint           string `protobuf:"bytes,4,opt,name=service_endpoint,json=serviceEndpoint,proto3" json:"service_endpoint,omitempty"`
	ServiceEndpointInsecure   bool   `protobuf:"varint,5,opt,name=service_endpoint_insecure,json=serviceEndpointInsecure,proto3" json:"service_endpoint_insecure,omitempty"`
	ServiceEncryptionKey      []byte `protobuf:"bytes,6,opt,name=service_encryption_key,json=serviceEncryptionKey,proto3" json:"service_encryption_key,omitempty"`
	ServiceClusterId          string `protobuf:"bytes,7,opt,name=service_cluster_id,json=serviceClusterId,proto3" json:"service_cluster_id,omitempty"`
}

func (x *ConfigSpec) Reset() {
	*x = ConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSpec) ProtoMessage() {}

func (x *ConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSpec.ProtoReflect.Descriptor instead.
func (*ConfigSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_cluster_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigSpec) GetDiscoveryEnabled() bool {
	if x != nil {
		return x.DiscoveryEnabled
	}
	return false
}

func (x *ConfigSpec) GetRegistryKubernetesEnabled() bool {
	if x != nil {
		return x.RegistryKubernetesEnabled
	}
	return false
}

func (x *ConfigSpec) GetRegistryServiceEnabled() bool {
	if x != nil {
		return x.RegistryServiceEnabled
	}
	return false
}

func (x *ConfigSpec) GetServiceEndpoint() string {
	if x != nil {
		return x.ServiceEndpoint
	}
	return ""
}

func (x *ConfigSpec) GetServiceEndpointInsecure() bool {
	if x != nil {
		return x.ServiceEndpointInsecure
	}
	return false
}

func (x *ConfigSpec) GetServiceEncryptionKey() []byte {
	if x != nil {
		return x.ServiceEncryptionKey
	}
	return nil
}

func (x *ConfigSpec) GetServiceClusterId() string {
	if x != nil {
		return x.ServiceClusterId
	}
	return ""
}

// ControlPlane describes ControlPlane data if any.
type ControlPlane struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiServerPort int64 `protobuf:"varint,1,opt,name=api_server_port,json=apiServerPort,proto3" json:"api_server_port,omitempty"`
}

func (x *ControlPlane) Reset() {
	*x = ControlPlane{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlPlane) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPlane) ProtoMessage() {}

func (x *ControlPlane) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPlane.ProtoReflect.Descriptor instead.
func (*ControlPlane) Descriptor() ([]byte, []int) {
	return file_resource_definitions_cluster_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *ControlPlane) GetApiServerPort() int64 {
	if x != nil {
		return x.ApiServerPort
	}
	return 0
}

// IdentitySpec describes status of rendered secrets.
//
// Note: IdentitySpec is persisted on disk in the STATE partition,
// so YAML serialization should be kept backwards compatible.
type IdentitySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *IdentitySpec) Reset() {
	*x = IdentitySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentitySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentitySpec) ProtoMessage() {}

func (x *IdentitySpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentitySpec.ProtoReflect.Descriptor instead.
func (*IdentitySpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_cluster_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *IdentitySpec) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

// InfoSpec describes cluster information.
type InfoSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId   string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ClusterName string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (x *InfoSpec) Reset() {
	*x = InfoSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoSpec) ProtoMessage() {}

func (x *InfoSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoSpec.ProtoReflect.Descriptor instead.
func (*InfoSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_cluster_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *InfoSpec) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *InfoSpec) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

// KubeSpanAffiliateSpec describes additional information specific for the KubeSpan.
type KubeSpanAffiliateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey           string                `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Address             *common.NetIP         `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	AdditionalAddresses []*common.NetIPPrefix `protobuf:"bytes,3,rep,name=additional_addresses,json=additionalAddresses,proto3" json:"additional_addresses,omitempty"`
	Endpoints           []*common.NetIPPort   `protobuf:"bytes,4,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *KubeSpanAffiliateSpec) Reset() {
	*x = KubeSpanAffiliateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeSpanAffiliateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeSpanAffiliateSpec) ProtoMessage() {}

func (x *KubeSpanAffiliateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeSpanAffiliateSpec.ProtoReflect.Descriptor instead.
func (*KubeSpanAffiliateSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_cluster_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *KubeSpanAffiliateSpec) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *KubeSpanAffiliateSpec) GetAddress() *common.NetIP {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *KubeSpanAffiliateSpec) GetAdditionalAddresses() []*common.NetIPPrefix {
	if x != nil {
		return x.AdditionalAddresses
	}
	return nil
}

func (x *KubeSpanAffiliateSpec) GetEndpoints() []*common.NetIPPort {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

// MemberSpec describes Member state.
type MemberSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId          string            `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Addresses       []*common.NetIP   `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Hostname        string            `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	MachineType     enums.MachineType `protobuf:"varint,4,opt,name=machine_type,json=machineType,proto3,enum=talos.resource.definitions.enums.MachineType" json:"machine_type,omitempty"`
	OperatingSystem string            `protobuf:"bytes,5,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
	ControlPlane    *ControlPlane     `protobuf:"bytes,6,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
}

func (x *MemberSpec) Reset() {
	*x = MemberSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberSpec) ProtoMessage() {}

func (x *MemberSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_cluster_cluster_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberSpec.ProtoReflect.Descriptor instead.
func (*MemberSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_cluster_cluster_proto_rawDescGZIP(), []int{6}
}

func (x *MemberSpec) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *MemberSpec) GetAddresses() []*common.NetIP {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *MemberSpec) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *MemberSpec) GetMachineType() enums.MachineType {
	if x != nil {
		return x.MachineType
	}
	return enums.MachineType(0)
}

func (x *MemberSpec) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *MemberSpec) GetControlPlane() *ControlPlane {
	if x != nil {
		return x.ControlPlane
	}
	return nil
}

var File_resource_definitions_cluster_cluster_proto protoreflect.FileDescriptor

var file_resource_definitions_cluster_cluster_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x74, 0x61,
	0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x03,
	0x0a, 0x0d, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x50, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d,
	0x2e, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x56, 0x0a, 0x09, 0x6b, 0x75,
	0x62, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x41, 0x66, 0x66, 0x69, 0x6c,
	0x69, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x6b, 0x75, 0x62, 0x65, 0x53, 0x70,
	0x61, 0x6e, 0x12, 0x55, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x61, 0x6c, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x22, 0xfe, 0x02, 0x0a, 0x0a, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x1b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x19, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x12,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x0c, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f,
	0x72, 0x74, 0x22, 0x27, 0x0a, 0x0c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x08, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x15, 0x4b, 0x75,
	0x62, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74,
	0x49, 0x50, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x14, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x13,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4e, 0x65, 0x74, 0x49, 0x50, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x22, 0xc2, 0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x52, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x74, 0x61,
	0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x55, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x61, 0x6c, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_definitions_cluster_cluster_proto_rawDescOnce sync.Once
	file_resource_definitions_cluster_cluster_proto_rawDescData = file_resource_definitions_cluster_cluster_proto_rawDesc
)

func file_resource_definitions_cluster_cluster_proto_rawDescGZIP() []byte {
	file_resource_definitions_cluster_cluster_proto_rawDescOnce.Do(func() {
		file_resource_definitions_cluster_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_definitions_cluster_cluster_proto_rawDescData)
	})
	return file_resource_definitions_cluster_cluster_proto_rawDescData
}

var file_resource_definitions_cluster_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_resource_definitions_cluster_cluster_proto_goTypes = []any{
	(*AffiliateSpec)(nil),         // 0: talos.resource.definitions.cluster.AffiliateSpec
	(*ConfigSpec)(nil),            // 1: talos.resource.definitions.cluster.ConfigSpec
	(*ControlPlane)(nil),          // 2: talos.resource.definitions.cluster.ControlPlane
	(*IdentitySpec)(nil),          // 3: talos.resource.definitions.cluster.IdentitySpec
	(*InfoSpec)(nil),              // 4: talos.resource.definitions.cluster.InfoSpec
	(*KubeSpanAffiliateSpec)(nil), // 5: talos.resource.definitions.cluster.KubeSpanAffiliateSpec
	(*MemberSpec)(nil),            // 6: talos.resource.definitions.cluster.MemberSpec
	(*common.NetIP)(nil),          // 7: common.NetIP
	(enums.MachineType)(0),        // 8: talos.resource.definitions.enums.MachineType
	(*common.NetIPPrefix)(nil),    // 9: common.NetIPPrefix
	(*common.NetIPPort)(nil),      // 10: common.NetIPPort
}
var file_resource_definitions_cluster_cluster_proto_depIdxs = []int32{
	7,  // 0: talos.resource.definitions.cluster.AffiliateSpec.addresses:type_name -> common.NetIP
	8,  // 1: talos.resource.definitions.cluster.AffiliateSpec.machine_type:type_name -> talos.resource.definitions.enums.MachineType
	5,  // 2: talos.resource.definitions.cluster.AffiliateSpec.kube_span:type_name -> talos.resource.definitions.cluster.KubeSpanAffiliateSpec
	2,  // 3: talos.resource.definitions.cluster.AffiliateSpec.control_plane:type_name -> talos.resource.definitions.cluster.ControlPlane
	7,  // 4: talos.resource.definitions.cluster.KubeSpanAffiliateSpec.address:type_name -> common.NetIP
	9,  // 5: talos.resource.definitions.cluster.KubeSpanAffiliateSpec.additional_addresses:type_name -> common.NetIPPrefix
	10, // 6: talos.resource.definitions.cluster.KubeSpanAffiliateSpec.endpoints:type_name -> common.NetIPPort
	7,  // 7: talos.resource.definitions.cluster.MemberSpec.addresses:type_name -> common.NetIP
	8,  // 8: talos.resource.definitions.cluster.MemberSpec.machine_type:type_name -> talos.resource.definitions.enums.MachineType
	2,  // 9: talos.resource.definitions.cluster.MemberSpec.control_plane:type_name -> talos.resource.definitions.cluster.ControlPlane
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_resource_definitions_cluster_cluster_proto_init() }
func file_resource_definitions_cluster_cluster_proto_init() {
	if File_resource_definitions_cluster_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_definitions_cluster_cluster_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AffiliateSpec); i {
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
		file_resource_definitions_cluster_cluster_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ConfigSpec); i {
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
		file_resource_definitions_cluster_cluster_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ControlPlane); i {
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
		file_resource_definitions_cluster_cluster_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IdentitySpec); i {
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
		file_resource_definitions_cluster_cluster_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*InfoSpec); i {
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
		file_resource_definitions_cluster_cluster_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*KubeSpanAffiliateSpec); i {
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
		file_resource_definitions_cluster_cluster_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*MemberSpec); i {
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
			RawDescriptor: file_resource_definitions_cluster_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_cluster_cluster_proto_goTypes,
		DependencyIndexes: file_resource_definitions_cluster_cluster_proto_depIdxs,
		MessageInfos:      file_resource_definitions_cluster_cluster_proto_msgTypes,
	}.Build()
	File_resource_definitions_cluster_cluster_proto = out.File
	file_resource_definitions_cluster_cluster_proto_rawDesc = nil
	file_resource_definitions_cluster_cluster_proto_goTypes = nil
	file_resource_definitions_cluster_cluster_proto_depIdxs = nil
}
