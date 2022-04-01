// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: v1/observability/observability.proto

package observability

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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request       string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	ClusterName   string `protobuf:"bytes,2,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Namespace     string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ContainerName string `protobuf:"bytes,4,opt,name=containerName,proto3" json:"containerName,omitempty"`
	Labels        string `protobuf:"bytes,5,opt,name=labels,proto3" json:"labels,omitempty"`
	FromSource    string `protobuf:"bytes,6,opt,name=fromSource,proto3" json:"fromSource,omitempty"`
	Duration      string `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *Data) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Data) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Data) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *Data) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

func (x *Data) GetFromSource() string {
	if x != nil {
		return x.FromSource
	}
	return ""
}

func (x *Data) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

type SysObsResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName   string                `protobuf:"bytes,1,opt,name=ClusterName,proto3" json:"ClusterName,omitempty"`
	NameSpace     string                `protobuf:"bytes,2,opt,name=NameSpace,proto3" json:"NameSpace,omitempty"`
	Labels        string                `protobuf:"bytes,3,opt,name=Labels,proto3" json:"Labels,omitempty"`
	ContainerName string                `protobuf:"bytes,4,opt,name=ContainerName,proto3" json:"ContainerName,omitempty"`
	Resources     []*SysProcessFileData `protobuf:"bytes,5,rep,name=Resources,proto3" json:"Resources,omitempty"`
}

func (x *SysObsResponseData) Reset() {
	*x = SysObsResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysObsResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysObsResponseData) ProtoMessage() {}

func (x *SysObsResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysObsResponseData.ProtoReflect.Descriptor instead.
func (*SysObsResponseData) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{1}
}

func (x *SysObsResponseData) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *SysObsResponseData) GetNameSpace() string {
	if x != nil {
		return x.NameSpace
	}
	return ""
}

func (x *SysObsResponseData) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

func (x *SysObsResponseData) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *SysObsResponseData) GetResources() []*SysProcessFileData {
	if x != nil {
		return x.Resources
	}
	return nil
}

type SysProcessFileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromSource   string   `protobuf:"bytes,1,opt,name=fromSource,proto3" json:"fromSource,omitempty"`
	ProcessPaths []string `protobuf:"bytes,2,rep,name=processPaths,proto3" json:"processPaths,omitempty"`
	FilePaths    []string `protobuf:"bytes,3,rep,name=filePaths,proto3" json:"filePaths,omitempty"`
	NetworkPaths []string `protobuf:"bytes,4,rep,name=networkPaths,proto3" json:"networkPaths,omitempty"`
}

func (x *SysProcessFileData) Reset() {
	*x = SysProcessFileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysProcessFileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysProcessFileData) ProtoMessage() {}

func (x *SysProcessFileData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysProcessFileData.ProtoReflect.Descriptor instead.
func (*SysProcessFileData) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{2}
}

func (x *SysProcessFileData) GetFromSource() string {
	if x != nil {
		return x.FromSource
	}
	return ""
}

func (x *SysProcessFileData) GetProcessPaths() []string {
	if x != nil {
		return x.ProcessPaths
	}
	return nil
}

func (x *SysProcessFileData) GetFilePaths() []string {
	if x != nil {
		return x.FilePaths
	}
	return nil
}

func (x *SysProcessFileData) GetNetworkPaths() []string {
	if x != nil {
		return x.NetworkPaths
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*SysObsResponseData `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_observability_observability_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_v1_observability_observability_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_v1_observability_observability_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetData() []*SysObsResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_v1_observability_observability_proto protoreflect.FileDescriptor

var file_v1_observability_observability_proto_rawDesc = []byte{
	0x0a, 0x24, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xda, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x6f,
	0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66,
	0x72, 0x6f, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x12, 0x53, 0x79, 0x73, 0x4f, 0x62, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x53, 0x79, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x9a,
	0x01, 0x0a, 0x12, 0x53, 0x79, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x50, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x73, 0x22, 0x44, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x79, 0x73, 0x4f, 0x62, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x32, 0x5b, 0x0a, 0x0d, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x4a, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x63,
	0x75, 0x6b, 0x6e, 0x6f, 0x78, 0x2f, 0x6b, 0x6e, 0x6f, 0x78, 0x41, 0x75, 0x74, 0x6f, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_observability_observability_proto_rawDescOnce sync.Once
	file_v1_observability_observability_proto_rawDescData = file_v1_observability_observability_proto_rawDesc
)

func file_v1_observability_observability_proto_rawDescGZIP() []byte {
	file_v1_observability_observability_proto_rawDescOnce.Do(func() {
		file_v1_observability_observability_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_observability_observability_proto_rawDescData)
	})
	return file_v1_observability_observability_proto_rawDescData
}

var file_v1_observability_observability_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_observability_observability_proto_goTypes = []interface{}{
	(*Data)(nil),               // 0: v1.observability.Data
	(*SysObsResponseData)(nil), // 1: v1.observability.SysObsResponseData
	(*SysProcessFileData)(nil), // 2: v1.observability.SysProcessFileData
	(*Response)(nil),           // 3: v1.observability.Response
}
var file_v1_observability_observability_proto_depIdxs = []int32{
	2, // 0: v1.observability.SysObsResponseData.Resources:type_name -> v1.observability.SysProcessFileData
	1, // 1: v1.observability.Response.Data:type_name -> v1.observability.SysObsResponseData
	0, // 2: v1.observability.Observability.SysObservabilityData:input_type -> v1.observability.Data
	3, // 3: v1.observability.Observability.SysObservabilityData:output_type -> v1.observability.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_observability_observability_proto_init() }
func file_v1_observability_observability_proto_init() {
	if File_v1_observability_observability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_observability_observability_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_v1_observability_observability_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysObsResponseData); i {
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
		file_v1_observability_observability_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysProcessFileData); i {
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
		file_v1_observability_observability_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_v1_observability_observability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_observability_observability_proto_goTypes,
		DependencyIndexes: file_v1_observability_observability_proto_depIdxs,
		MessageInfos:      file_v1_observability_observability_proto_msgTypes,
	}.Build()
	File_v1_observability_observability_proto = out.File
	file_v1_observability_observability_proto_rawDesc = nil
	file_v1_observability_observability_proto_goTypes = nil
	file_v1_observability_observability_proto_depIdxs = nil
}
