// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/compression/gzip/decompressor/v3/gzip.proto

package envoy_extensions_compression_gzip_decompressor_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Gzip struct {
	WindowBits           *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	ChunkSize            *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Gzip) Reset()         { *m = Gzip{} }
func (m *Gzip) String() string { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()    {}
func (*Gzip) Descriptor() ([]byte, []int) {
	return fileDescriptor_6de1f95e053796e8, []int{0}
}

func (m *Gzip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip.Unmarshal(m, b)
}
func (m *Gzip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip.Marshal(b, m, deterministic)
}
func (m *Gzip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip.Merge(m, src)
}
func (m *Gzip) XXX_Size() int {
	return xxx_messageInfo_Gzip.Size(m)
}
func (m *Gzip) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip proto.InternalMessageInfo

func (m *Gzip) GetWindowBits() *wrappers.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

func (m *Gzip) GetChunkSize() *wrappers.UInt32Value {
	if m != nil {
		return m.ChunkSize
	}
	return nil
}

func init() {
	proto.RegisterType((*Gzip)(nil), "envoy.extensions.compression.gzip.decompressor.v3.Gzip")
}

func init() {
	proto.RegisterFile("envoy/extensions/compression/gzip/decompressor/v3/gzip.proto", fileDescriptor_6de1f95e053796e8)
}

var fileDescriptor_6de1f95e053796e8 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x6e, 0xf2, 0x30,
	0x10, 0xc7, 0x15, 0x84, 0xf8, 0x14, 0xf3, 0x49, 0xad, 0x58, 0x1a, 0xa1, 0x16, 0xd1, 0x4e, 0x88,
	0xc1, 0x56, 0xc9, 0x5a, 0xa9, 0x52, 0x96, 0x96, 0x0d, 0x51, 0x95, 0x15, 0x19, 0x72, 0x4d, 0xad,
	0x82, 0xcf, 0xb2, 0x9d, 0x00, 0x99, 0xf2, 0x1e, 0x7d, 0x93, 0x3e, 0x41, 0xd7, 0xbe, 0x0e, 0x53,
	0x15, 0x87, 0xa8, 0x48, 0x2c, 0xed, 0x66, 0xdf, 0xff, 0x7e, 0xe7, 0x9f, 0x6d, 0x72, 0x07, 0x32,
	0xc3, 0x1d, 0x83, 0xad, 0x05, 0x69, 0x04, 0x4a, 0xc3, 0x96, 0xb8, 0x56, 0x1a, 0x4c, 0xb9, 0x61,
	0x49, 0x2e, 0x14, 0x8b, 0xa1, 0x2e, 0xa1, 0x66, 0x59, 0xe8, 0x8a, 0x54, 0x69, 0xb4, 0xd8, 0xb9,
	0x75, 0x34, 0xfd, 0xa1, 0xe9, 0x11, 0x4d, 0x5d, 0xe3, 0x31, 0x4d, 0xb3, 0xb0, 0xdb, 0x4b, 0x10,
	0x93, 0x15, 0x30, 0x37, 0x60, 0x91, 0xbe, 0xb0, 0x8d, 0xe6, 0x4a, 0x81, 0x36, 0xd5, 0xc8, 0x6e,
	0x2f, 0x8d, 0x15, 0x67, 0x5c, 0x4a, 0xb4, 0xdc, 0x3a, 0xa1, 0xb5, 0x48, 0x34, 0xb7, 0x70, 0xc8,
	0xaf, 0x4e, 0x72, 0x63, 0xb9, 0x4d, 0x6b, 0xfc, 0xfa, 0x24, 0xce, 0x40, 0x97, 0x36, 0x42, 0x26,
	0x87, 0x96, 0x8b, 0x8c, 0xaf, 0x44, 0xcc, 0x2d, 0xb0, 0x7a, 0x51, 0x05, 0x37, 0xef, 0x1e, 0x69,
	0x3e, 0xe4, 0x42, 0x75, 0x1e, 0x49, 0x7b, 0x23, 0x64, 0x8c, 0x9b, 0xf9, 0x42, 0x58, 0x13, 0x78,
	0x7d, 0x6f, 0xd0, 0x1e, 0x5d, 0xd2, 0xca, 0x9c, 0xd6, 0xe6, 0xf4, 0x79, 0x2c, 0x6d, 0x38, 0x9a,
	0xf1, 0x55, 0x0a, 0x91, 0xbf, 0x8f, 0x5a, 0xc3, 0x66, 0x70, 0x36, 0xf0, 0xa7, 0xa4, 0x62, 0x23,
	0x61, 0x4d, 0x67, 0x4c, 0xc8, 0xf2, 0x35, 0x95, 0x6f, 0x73, 0x23, 0x72, 0x08, 0x1a, 0xbf, 0x18,
	0xf4, 0x7f, 0x1f, 0xf9, 0xc3, 0x7f, 0x41, 0x51, 0x34, 0x07, 0x45, 0x7f, 0xea, 0x3b, 0xfa, 0x49,
	0xe4, 0x10, 0xcd, 0x3e, 0x8a, 0xcf, 0xaf, 0x56, 0xe3, 0xbc, 0x41, 0xee, 0x05, 0x52, 0xf7, 0xf0,
	0x4a, 0xe3, 0x76, 0x47, 0xff, 0xfc, 0x07, 0x91, 0x5f, 0xde, 0x72, 0x52, 0x9e, 0x3e, 0xf1, 0x16,
	0x2d, 0xa7, 0x11, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x43, 0x4d, 0x10, 0x0a, 0x02, 0x00,
	0x00,
}
