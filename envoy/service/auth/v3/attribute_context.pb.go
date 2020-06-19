// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/service/auth/v3/attribute_context.proto

package envoy_service_auth_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AttributeContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source            *AttributeContext_Peer    `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination       *AttributeContext_Peer    `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Request           *AttributeContext_Request `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	ContextExtensions map[string]string         `protobuf:"bytes,10,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MetadataContext   *v3.Metadata              `protobuf:"bytes,11,opt,name=metadata_context,json=metadataContext,proto3" json:"metadata_context,omitempty"`
}

func (x *AttributeContext) Reset() {
	*x = AttributeContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_auth_v3_attribute_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeContext) ProtoMessage() {}

func (x *AttributeContext) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_auth_v3_attribute_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeContext.ProtoReflect.Descriptor instead.
func (*AttributeContext) Descriptor() ([]byte, []int) {
	return file_envoy_service_auth_v3_attribute_context_proto_rawDescGZIP(), []int{0}
}

func (x *AttributeContext) GetSource() *AttributeContext_Peer {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *AttributeContext) GetDestination() *AttributeContext_Peer {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *AttributeContext) GetRequest() *AttributeContext_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *AttributeContext) GetContextExtensions() map[string]string {
	if x != nil {
		return x.ContextExtensions
	}
	return nil
}

func (x *AttributeContext) GetMetadataContext() *v3.Metadata {
	if x != nil {
		return x.MetadataContext
	}
	return nil
}

type AttributeContext_Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     *v3.Address       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Service     string            `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Labels      map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Principal   string            `protobuf:"bytes,4,opt,name=principal,proto3" json:"principal,omitempty"`
	Certificate string            `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *AttributeContext_Peer) Reset() {
	*x = AttributeContext_Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_auth_v3_attribute_context_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeContext_Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeContext_Peer) ProtoMessage() {}

func (x *AttributeContext_Peer) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_auth_v3_attribute_context_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeContext_Peer.ProtoReflect.Descriptor instead.
func (*AttributeContext_Peer) Descriptor() ([]byte, []int) {
	return file_envoy_service_auth_v3_attribute_context_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AttributeContext_Peer) GetAddress() *v3.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *AttributeContext_Peer) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AttributeContext_Peer) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AttributeContext_Peer) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *AttributeContext_Peer) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

type AttributeContext_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamp.Timestamp          `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Http *AttributeContext_HttpRequest `protobuf:"bytes,2,opt,name=http,proto3" json:"http,omitempty"`
}

func (x *AttributeContext_Request) Reset() {
	*x = AttributeContext_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_auth_v3_attribute_context_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeContext_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeContext_Request) ProtoMessage() {}

func (x *AttributeContext_Request) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_auth_v3_attribute_context_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeContext_Request.ProtoReflect.Descriptor instead.
func (*AttributeContext_Request) Descriptor() ([]byte, []int) {
	return file_envoy_service_auth_v3_attribute_context_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AttributeContext_Request) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *AttributeContext_Request) GetHttp() *AttributeContext_HttpRequest {
	if x != nil {
		return x.Http
	}
	return nil
}

type AttributeContext_HttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Method   string            `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Headers  map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Path     string            `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Host     string            `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Scheme   string            `protobuf:"bytes,6,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Query    string            `protobuf:"bytes,7,opt,name=query,proto3" json:"query,omitempty"`
	Fragment string            `protobuf:"bytes,8,opt,name=fragment,proto3" json:"fragment,omitempty"`
	Size     int64             `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Protocol string            `protobuf:"bytes,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Body     string            `protobuf:"bytes,11,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *AttributeContext_HttpRequest) Reset() {
	*x = AttributeContext_HttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_auth_v3_attribute_context_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeContext_HttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeContext_HttpRequest) ProtoMessage() {}

func (x *AttributeContext_HttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_auth_v3_attribute_context_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeContext_HttpRequest.ProtoReflect.Descriptor instead.
func (*AttributeContext_HttpRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_auth_v3_attribute_context_proto_rawDescGZIP(), []int{0, 2}
}

func (x *AttributeContext_HttpRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AttributeContext_HttpRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *AttributeContext_HttpRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *AttributeContext_HttpRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AttributeContext_HttpRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AttributeContext_HttpRequest) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *AttributeContext_HttpRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *AttributeContext_HttpRequest) GetFragment() string {
	if x != nil {
		return x.Fragment
	}
	return ""
}

func (x *AttributeContext_HttpRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *AttributeContext_HttpRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *AttributeContext_HttpRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_envoy_service_auth_v3_attribute_context_proto protoreflect.FileDescriptor

var file_envoy_service_auth_v3_attribute_context_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc,
	0x0b, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x33, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x6d, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x49, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0xda,
	0x02, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x33, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x32, 0x9a, 0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x1a, 0xb9, 0x01, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70,
	0x3a, 0x35, 0x9a, 0xc5, 0x88, 0x1e, 0x30, 0x0a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0xbe, 0x03, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x5a, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x40, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x39, 0x9a,
	0xc5, 0x88, 0x1e, 0x34, 0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x2d,
	0x9a, 0xc5, 0x88, 0x1e, 0x28, 0x0a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x46, 0x0a,
	0x23, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x33, 0x42, 0x15, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_auth_v3_attribute_context_proto_rawDescOnce sync.Once
	file_envoy_service_auth_v3_attribute_context_proto_rawDescData = file_envoy_service_auth_v3_attribute_context_proto_rawDesc
)

func file_envoy_service_auth_v3_attribute_context_proto_rawDescGZIP() []byte {
	file_envoy_service_auth_v3_attribute_context_proto_rawDescOnce.Do(func() {
		file_envoy_service_auth_v3_attribute_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_auth_v3_attribute_context_proto_rawDescData)
	})
	return file_envoy_service_auth_v3_attribute_context_proto_rawDescData
}

var file_envoy_service_auth_v3_attribute_context_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_envoy_service_auth_v3_attribute_context_proto_goTypes = []interface{}{
	(*AttributeContext)(nil),             // 0: envoy.service.auth.v3.AttributeContext
	(*AttributeContext_Peer)(nil),        // 1: envoy.service.auth.v3.AttributeContext.Peer
	(*AttributeContext_Request)(nil),     // 2: envoy.service.auth.v3.AttributeContext.Request
	(*AttributeContext_HttpRequest)(nil), // 3: envoy.service.auth.v3.AttributeContext.HttpRequest
	nil,                                  // 4: envoy.service.auth.v3.AttributeContext.ContextExtensionsEntry
	nil,                                  // 5: envoy.service.auth.v3.AttributeContext.Peer.LabelsEntry
	nil,                                  // 6: envoy.service.auth.v3.AttributeContext.HttpRequest.HeadersEntry
	(*v3.Metadata)(nil),                  // 7: envoy.config.core.v3.Metadata
	(*v3.Address)(nil),                   // 8: envoy.config.core.v3.Address
	(*timestamp.Timestamp)(nil),          // 9: google.protobuf.Timestamp
}
var file_envoy_service_auth_v3_attribute_context_proto_depIdxs = []int32{
	1,  // 0: envoy.service.auth.v3.AttributeContext.source:type_name -> envoy.service.auth.v3.AttributeContext.Peer
	1,  // 1: envoy.service.auth.v3.AttributeContext.destination:type_name -> envoy.service.auth.v3.AttributeContext.Peer
	2,  // 2: envoy.service.auth.v3.AttributeContext.request:type_name -> envoy.service.auth.v3.AttributeContext.Request
	4,  // 3: envoy.service.auth.v3.AttributeContext.context_extensions:type_name -> envoy.service.auth.v3.AttributeContext.ContextExtensionsEntry
	7,  // 4: envoy.service.auth.v3.AttributeContext.metadata_context:type_name -> envoy.config.core.v3.Metadata
	8,  // 5: envoy.service.auth.v3.AttributeContext.Peer.address:type_name -> envoy.config.core.v3.Address
	5,  // 6: envoy.service.auth.v3.AttributeContext.Peer.labels:type_name -> envoy.service.auth.v3.AttributeContext.Peer.LabelsEntry
	9,  // 7: envoy.service.auth.v3.AttributeContext.Request.time:type_name -> google.protobuf.Timestamp
	3,  // 8: envoy.service.auth.v3.AttributeContext.Request.http:type_name -> envoy.service.auth.v3.AttributeContext.HttpRequest
	6,  // 9: envoy.service.auth.v3.AttributeContext.HttpRequest.headers:type_name -> envoy.service.auth.v3.AttributeContext.HttpRequest.HeadersEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_envoy_service_auth_v3_attribute_context_proto_init() }
func file_envoy_service_auth_v3_attribute_context_proto_init() {
	if File_envoy_service_auth_v3_attribute_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_auth_v3_attribute_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeContext); i {
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
		file_envoy_service_auth_v3_attribute_context_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeContext_Peer); i {
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
		file_envoy_service_auth_v3_attribute_context_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeContext_Request); i {
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
		file_envoy_service_auth_v3_attribute_context_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeContext_HttpRequest); i {
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
			RawDescriptor: file_envoy_service_auth_v3_attribute_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_service_auth_v3_attribute_context_proto_goTypes,
		DependencyIndexes: file_envoy_service_auth_v3_attribute_context_proto_depIdxs,
		MessageInfos:      file_envoy_service_auth_v3_attribute_context_proto_msgTypes,
	}.Build()
	File_envoy_service_auth_v3_attribute_context_proto = out.File
	file_envoy_service_auth_v3_attribute_context_proto_rawDesc = nil
	file_envoy_service_auth_v3_attribute_context_proto_goTypes = nil
	file_envoy_service_auth_v3_attribute_context_proto_depIdxs = nil
}