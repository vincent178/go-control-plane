// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/compressor/v2/compressor.proto

package envoy_config_filter_http_compressor_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
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

type Compressor struct {
	ContentLength              *wrappers.UInt32Value    `protobuf:"bytes,1,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	ContentType                []string                 `protobuf:"bytes,2,rep,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	DisableOnEtagHeader        bool                     `protobuf:"varint,3,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	RemoveAcceptEncodingHeader bool                     `protobuf:"varint,4,opt,name=remove_accept_encoding_header,json=removeAcceptEncodingHeader,proto3" json:"remove_accept_encoding_header,omitempty"`
	RuntimeEnabled             *core.RuntimeFeatureFlag `protobuf:"bytes,5,opt,name=runtime_enabled,json=runtimeEnabled,proto3" json:"runtime_enabled,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                 `json:"-"`
	XXX_unrecognized           []byte                   `json:"-"`
	XXX_sizecache              int32                    `json:"-"`
}

func (m *Compressor) Reset()         { *m = Compressor{} }
func (m *Compressor) String() string { return proto.CompactTextString(m) }
func (*Compressor) ProtoMessage()    {}
func (*Compressor) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ef4714f5f18ab25, []int{0}
}

func (m *Compressor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Compressor.Unmarshal(m, b)
}
func (m *Compressor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Compressor.Marshal(b, m, deterministic)
}
func (m *Compressor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Compressor.Merge(m, src)
}
func (m *Compressor) XXX_Size() int {
	return xxx_messageInfo_Compressor.Size(m)
}
func (m *Compressor) XXX_DiscardUnknown() {
	xxx_messageInfo_Compressor.DiscardUnknown(m)
}

var xxx_messageInfo_Compressor proto.InternalMessageInfo

func (m *Compressor) GetContentLength() *wrappers.UInt32Value {
	if m != nil {
		return m.ContentLength
	}
	return nil
}

func (m *Compressor) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Compressor) GetDisableOnEtagHeader() bool {
	if m != nil {
		return m.DisableOnEtagHeader
	}
	return false
}

func (m *Compressor) GetRemoveAcceptEncodingHeader() bool {
	if m != nil {
		return m.RemoveAcceptEncodingHeader
	}
	return false
}

func (m *Compressor) GetRuntimeEnabled() *core.RuntimeFeatureFlag {
	if m != nil {
		return m.RuntimeEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*Compressor)(nil), "envoy.config.filter.http.compressor.v2.Compressor")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/compressor/v2/compressor.proto", fileDescriptor_8ef4714f5f18ab25)
}

var fileDescriptor_8ef4714f5f18ab25 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0xd9, 0xd6, 0x8a, 0x4e, 0xb5, 0x85, 0x08, 0xb2, 0x2c, 0x55, 0x56, 0x41, 0x59, 0x10,
	0x67, 0x20, 0x2b, 0x78, 0xdd, 0x96, 0x2d, 0x0a, 0xa2, 0x25, 0xa8, 0xb7, 0xe1, 0x6c, 0x72, 0x36,
	0x3b, 0x90, 0x3d, 0x67, 0x98, 0x9c, 0xc4, 0x2e, 0xf8, 0x0e, 0xde, 0xfa, 0x30, 0x3e, 0x95, 0x0f,
	0x20, 0x92, 0xcc, 0x6c, 0x2d, 0x78, 0xd3, 0xcb, 0x93, 0xf3, 0x7f, 0xff, 0x4f, 0xce, 0x3f, 0xea,
	0x2d, 0x52, 0xc7, 0x5b, 0x53, 0x30, 0xad, 0x6c, 0x65, 0x56, 0xb6, 0x16, 0xf4, 0x66, 0x2d, 0xe2,
	0x4c, 0xc1, 0x1b, 0xe7, 0xb1, 0x69, 0xd8, 0x9b, 0x2e, 0xbd, 0x31, 0x69, 0xe7, 0x59, 0x38, 0x79,
	0x39, 0x80, 0x3a, 0x80, 0x3a, 0x80, 0xba, 0x07, 0xf5, 0x0d, 0x69, 0x97, 0x4e, 0x4e, 0x42, 0x00,
	0x38, 0x1b, 0x6c, 0x3c, 0x9a, 0x25, 0x34, 0x18, 0x5c, 0x26, 0x4f, 0x2b, 0xe6, 0xaa, 0x46, 0x33,
	0x4c, 0xcb, 0x76, 0x65, 0xbe, 0x79, 0x70, 0x0e, 0x7d, 0xb3, 0xdb, 0xb7, 0xa5, 0x03, 0x03, 0x44,
	0x2c, 0x20, 0x96, 0xa9, 0x31, 0x1b, 0x5b, 0x79, 0x90, 0xc8, 0x3f, 0xff, 0xb5, 0xa7, 0xd4, 0xf9,
	0x75, 0x5e, 0x72, 0xae, 0x8e, 0x0a, 0x26, 0x41, 0x92, 0xbc, 0x46, 0xaa, 0x64, 0x3d, 0x1e, 0x4d,
	0x47, 0xb3, 0xc3, 0xf4, 0x44, 0x87, 0x1c, 0xbd, 0xcb, 0xd1, 0x5f, 0xde, 0x93, 0xcc, 0xd3, 0xaf,
	0x50, 0xb7, 0x98, 0x3d, 0x8c, 0xcc, 0x87, 0x01, 0x49, 0x9e, 0xa9, 0x07, 0x3b, 0x13, 0xd9, 0x3a,
	0x1c, 0xef, 0x4d, 0xf7, 0x67, 0xf7, 0xb3, 0xc3, 0xf8, 0xed, 0xf3, 0xd6, 0x61, 0x32, 0x57, 0x8f,
	0x4b, 0xdb, 0xc0, 0xb2, 0xc6, 0x9c, 0x29, 0x47, 0x81, 0x2a, 0x5f, 0x23, 0x94, 0xe8, 0xc7, 0xfb,
	0xd3, 0xd1, 0xec, 0x5e, 0xf6, 0x28, 0x6e, 0x3f, 0xd1, 0x42, 0xa0, 0x7a, 0x37, 0xac, 0x92, 0x53,
	0xf5, 0xc4, 0xe3, 0x86, 0x3b, 0xcc, 0xa1, 0x28, 0xd0, 0x49, 0x8e, 0x54, 0x70, 0x69, 0xe9, 0x9a,
	0xbd, 0x33, 0xb0, 0x93, 0x20, 0x3a, 0x1d, 0x34, 0x8b, 0x28, 0x89, 0x16, 0x1f, 0xd5, 0xb1, 0x6f,
	0x49, 0xec, 0x06, 0x73, 0xa4, 0x3e, 0xa0, 0x1c, 0x1f, 0x0c, 0x3f, 0xf8, 0x42, 0x87, 0x3a, 0xc0,
	0x59, 0xdd, 0xa5, 0xba, 0x3f, 0xb3, 0xce, 0x82, 0xf2, 0x02, 0x41, 0x5a, 0x8f, 0x17, 0x35, 0x54,
	0xd9, 0x51, 0xa4, 0x17, 0x01, 0x3e, 0xfb, 0xfe, 0xfb, 0xe7, 0x9f, 0x1f, 0x07, 0xaf, 0x93, 0x57,
	0x81, 0xc6, 0x2b, 0x41, 0x6a, 0xfa, 0x33, 0xc7, 0x42, 0x9b, 0xff, 0x1b, 0x9d, 0xab, 0x37, 0x96,
	0x43, 0x9a, 0xf3, 0x7c, 0xb5, 0xd5, 0xb7, 0x7b, 0x07, 0x67, 0xc7, 0xff, 0x6a, 0xba, 0xec, 0x4b,
	0xb8, 0x1c, 0x2d, 0xef, 0x0e, 0x6d, 0xcc, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x59, 0xbc,
	0x0f, 0x84, 0x02, 0x00, 0x00,
}
