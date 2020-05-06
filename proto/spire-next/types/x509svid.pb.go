// Code generated by protoc-gen-go. DO NOT EDIT.
// source: x509svid.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// X.509 SPIFFE Verifiable Identity Document. It contains the raw X.509
// certificate data as well as a few denormalized fields for convenience.
type X509SVID struct {
	// Certificate and intermediates required to form a chain of trust back to
	// the X.509 authorities of the trust domain (ASN.1 DER encoded).
	CertChain [][]byte `protobuf:"bytes,1,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	// SPIFFE ID of the SVID.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Expiration timestamp (seconds since Unix epoch).
	ExpiresAt            int64    `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *X509SVID) Reset()         { *m = X509SVID{} }
func (m *X509SVID) String() string { return proto.CompactTextString(m) }
func (*X509SVID) ProtoMessage()    {}
func (*X509SVID) Descriptor() ([]byte, []int) {
	return fileDescriptor_876f77a3d30643db, []int{0}
}

func (m *X509SVID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_X509SVID.Unmarshal(m, b)
}
func (m *X509SVID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_X509SVID.Marshal(b, m, deterministic)
}
func (m *X509SVID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_X509SVID.Merge(m, src)
}
func (m *X509SVID) XXX_Size() int {
	return xxx_messageInfo_X509SVID.Size(m)
}
func (m *X509SVID) XXX_DiscardUnknown() {
	xxx_messageInfo_X509SVID.DiscardUnknown(m)
}

var xxx_messageInfo_X509SVID proto.InternalMessageInfo

func (m *X509SVID) GetCertChain() [][]byte {
	if m != nil {
		return m.CertChain
	}
	return nil
}

func (m *X509SVID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *X509SVID) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*X509SVID)(nil), "spire.types.X509SVID")
}

func init() { proto.RegisterFile("x509svid.proto", fileDescriptor_876f77a3d30643db) }

var fileDescriptor_876f77a3d30643db = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xab, 0x30, 0x35, 0xb0,
	0x2c, 0x2e, 0xcb, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x2e, 0xc8, 0x2c,
	0x4a, 0xd5, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x8a, 0xe0, 0xe2, 0x88, 0x30, 0x35, 0xb0, 0x0c,
	0x0e, 0xf3, 0x74, 0x11, 0x92, 0xe5, 0xe2, 0x4a, 0x4e, 0x2d, 0x2a, 0x89, 0x4f, 0xce, 0x48, 0xcc,
	0xcc, 0x93, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x09, 0xe2, 0x04, 0x89, 0x38, 0x83, 0x04, 0x84, 0xf8,
	0xb8, 0x98, 0x32, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x40, 0xca,
	0x53, 0x2b, 0x40, 0x46, 0x15, 0xc7, 0x27, 0x96, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0x30, 0x07, 0x71,
	0x42, 0x45, 0x1c, 0x4b, 0x9c, 0x0c, 0xa2, 0xf4, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0x8b, 0x0b, 0x32, 0xd3, 0xd2, 0x52, 0xf5, 0xc1, 0x56, 0xeb, 0x83, 0xdd, 0x01,
	0x61, 0xeb, 0xe6, 0xa5, 0x56, 0x94, 0xe8, 0x83, 0xdd, 0x92, 0xc4, 0x06, 0x16, 0x37, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0x36, 0x7f, 0x9f, 0xb1, 0x00, 0x00, 0x00,
}