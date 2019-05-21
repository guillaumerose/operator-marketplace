// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/geometry.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A vertex represents a 2D point in the image.
// The normalized vertex coordinates are between 0 to 1 fractions relative to
// the original plane (image, video). E.g. if the plane (e.g. whole image) would
// have size 10 x 20 then a point with normalized coordinates (0.1, 0.3) would
// be at the position (1, 6) on that plane.
type NormalizedVertex struct {
	// Required. Horizontal coordinate.
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	// Required. Vertical coordinate.
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NormalizedVertex) Reset()         { *m = NormalizedVertex{} }
func (m *NormalizedVertex) String() string { return proto.CompactTextString(m) }
func (*NormalizedVertex) ProtoMessage()    {}
func (*NormalizedVertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbaead6fc80bf3a2, []int{0}
}

func (m *NormalizedVertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormalizedVertex.Unmarshal(m, b)
}
func (m *NormalizedVertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormalizedVertex.Marshal(b, m, deterministic)
}
func (m *NormalizedVertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormalizedVertex.Merge(m, src)
}
func (m *NormalizedVertex) XXX_Size() int {
	return xxx_messageInfo_NormalizedVertex.Size(m)
}
func (m *NormalizedVertex) XXX_DiscardUnknown() {
	xxx_messageInfo_NormalizedVertex.DiscardUnknown(m)
}

var xxx_messageInfo_NormalizedVertex proto.InternalMessageInfo

func (m *NormalizedVertex) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *NormalizedVertex) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A bounding polygon of a detected object on a plane.
// On output both vertices and normalized_vertices are provided.
// The polygon is formed by connecting vertices in the order they are listed.
type BoundingPoly struct {
	// Output only . The bounding polygon normalized vertices.
	NormalizedVertices   []*NormalizedVertex `protobuf:"bytes,2,rep,name=normalized_vertices,json=normalizedVertices,proto3" json:"normalized_vertices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BoundingPoly) Reset()         { *m = BoundingPoly{} }
func (m *BoundingPoly) String() string { return proto.CompactTextString(m) }
func (*BoundingPoly) ProtoMessage()    {}
func (*BoundingPoly) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbaead6fc80bf3a2, []int{1}
}

func (m *BoundingPoly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingPoly.Unmarshal(m, b)
}
func (m *BoundingPoly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingPoly.Marshal(b, m, deterministic)
}
func (m *BoundingPoly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingPoly.Merge(m, src)
}
func (m *BoundingPoly) XXX_Size() int {
	return xxx_messageInfo_BoundingPoly.Size(m)
}
func (m *BoundingPoly) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingPoly.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingPoly proto.InternalMessageInfo

func (m *BoundingPoly) GetNormalizedVertices() []*NormalizedVertex {
	if m != nil {
		return m.NormalizedVertices
	}
	return nil
}

func init() {
	proto.RegisterType((*NormalizedVertex)(nil), "google.cloud.automl.v1beta1.NormalizedVertex")
	proto.RegisterType((*BoundingPoly)(nil), "google.cloud.automl.v1beta1.BoundingPoly")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/geometry.proto", fileDescriptor_cbaead6fc80bf3a2)
}

var fileDescriptor_cbaead6fc80bf3a2 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x49, 0x05, 0x0f, 0x71, 0x07, 0xa9, 0x97, 0xe2, 0x04, 0xc7, 0x4e, 0x43, 0x30, 0x61,
	0x7a, 0xf4, 0xb4, 0x79, 0xf0, 0xa4, 0x8c, 0x1d, 0x76, 0x90, 0xa2, 0x64, 0xed, 0x23, 0x04, 0xd2,
	0xf7, 0x4a, 0xfa, 0x3a, 0x56, 0xcf, 0x7e, 0x3a, 0x3f, 0x95, 0xac, 0x29, 0x82, 0x22, 0x3b, 0xfe,
	0xf3, 0x7e, 0xff, 0x5f, 0x92, 0x27, 0x6f, 0x2c, 0x91, 0xf5, 0xa0, 0x0b, 0x4f, 0x6d, 0xa9, 0x4d,
	0xcb, 0x54, 0x79, 0xbd, 0x9b, 0x6f, 0x81, 0xcd, 0x5c, 0x5b, 0xa0, 0x0a, 0x38, 0x74, 0xaa, 0x0e,
	0xc4, 0x94, 0x8e, 0x23, 0xab, 0x7a, 0x56, 0x45, 0x56, 0x0d, 0xec, 0xe5, 0xd5, 0x20, 0x32, 0xb5,
	0xd3, 0x06, 0x91, 0xd8, 0xb0, 0x23, 0x6c, 0x62, 0x75, 0xaa, 0xe4, 0xf9, 0x0b, 0x85, 0xca, 0x78,
	0xf7, 0x01, 0xe5, 0x06, 0x02, 0xc3, 0x3e, 0x1d, 0x49, 0xb1, 0xcf, 0xc4, 0x44, 0xcc, 0x92, 0xb5,
	0xe8, 0x53, 0x97, 0x25, 0x31, 0x75, 0x53, 0x94, 0xa3, 0x25, 0xb5, 0x58, 0x3a, 0xb4, 0x2b, 0xf2,
	0x5d, 0xfa, 0x26, 0x2f, 0xf0, 0xa7, 0xff, 0xbe, 0x83, 0xc0, 0xae, 0x80, 0x26, 0x4b, 0x26, 0x27,
	0xb3, 0xb3, 0xbb, 0x5b, 0x75, 0xe4, 0x61, 0xea, 0xef, 0xbd, 0xeb, 0x14, 0x7f, 0x9d, 0x1c, 0x44,
	0xcb, 0x4f, 0x21, 0xaf, 0x0b, 0xaa, 0x8e, 0x89, 0x56, 0xe2, 0x75, 0x31, 0x8c, 0x2d, 0x79, 0x83,
	0x56, 0x51, 0xb0, 0xda, 0x02, 0xf6, 0x3f, 0xd4, 0x71, 0x64, 0x6a, 0xd7, 0xfc, 0xbb, 0xcb, 0x87,
	0x18, 0xbf, 0x92, 0xf1, 0x53, 0x0f, 0xe6, 0x8f, 0x07, 0x28, 0x5f, 0xb4, 0x4c, 0xcf, 0x3e, 0xdf,
	0x44, 0x68, 0x7b, 0xda, 0xbb, 0xee, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x39, 0xfe, 0x55, 0x67,
	0x96, 0x01, 0x00, 0x00,
}