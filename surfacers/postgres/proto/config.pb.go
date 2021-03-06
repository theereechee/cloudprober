// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/surfacers/postgres/proto/config.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SurfacerConf struct {
	ConnectionString     *string  `protobuf:"bytes,1,req,name=connection_string,json=connectionString" json:"connection_string,omitempty"`
	MetricsTableName     *string  `protobuf:"bytes,2,req,name=metrics_table_name,json=metricsTableName" json:"metrics_table_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SurfacerConf) Reset()         { *m = SurfacerConf{} }
func (m *SurfacerConf) String() string { return proto.CompactTextString(m) }
func (*SurfacerConf) ProtoMessage()    {}
func (*SurfacerConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ca5f6a08327971c1, []int{0}
}
func (m *SurfacerConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurfacerConf.Unmarshal(m, b)
}
func (m *SurfacerConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurfacerConf.Marshal(b, m, deterministic)
}
func (dst *SurfacerConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurfacerConf.Merge(dst, src)
}
func (m *SurfacerConf) XXX_Size() int {
	return xxx_messageInfo_SurfacerConf.Size(m)
}
func (m *SurfacerConf) XXX_DiscardUnknown() {
	xxx_messageInfo_SurfacerConf.DiscardUnknown(m)
}

var xxx_messageInfo_SurfacerConf proto.InternalMessageInfo

func (m *SurfacerConf) GetConnectionString() string {
	if m != nil && m.ConnectionString != nil {
		return *m.ConnectionString
	}
	return ""
}

func (m *SurfacerConf) GetMetricsTableName() string {
	if m != nil && m.MetricsTableName != nil {
		return *m.MetricsTableName
	}
	return ""
}

func init() {
	proto.RegisterType((*SurfacerConf)(nil), "cloudprober.surfacer.postgres.SurfacerConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/surfacers/postgres/proto/config.proto", fileDescriptor_config_ca5f6a08327971c1)
}

var fileDescriptor_config_ca5f6a08327971c1 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8b, 0xb1, 0x8e, 0xc2, 0x30,
	0x10, 0x05, 0x75, 0xe9, 0xce, 0xba, 0xe2, 0xce, 0x55, 0x9a, 0x93, 0x10, 0x15, 0x12, 0xc8, 0xfe,
	0x88, 0xf4, 0x14, 0x84, 0x3e, 0x72, 0x96, 0x8d, 0xb1, 0x14, 0xef, 0x46, 0xeb, 0xcd, 0xff, 0x23,
	0x42, 0x22, 0x28, 0xdf, 0xcc, 0x1b, 0xd3, 0xc4, 0xa4, 0xf7, 0xb9, 0x77, 0xc0, 0xd9, 0x47, 0xe6,
	0x38, 0xa2, 0x87, 0x91, 0xe7, 0xdb, 0x24, 0xdc, 0xa3, 0xf8, 0x32, 0xcb, 0x10, 0x00, 0xa5, 0xf8,
	0x89, 0x8b, 0x46, 0xc1, 0xe2, 0x27, 0x61, 0x65, 0x0f, 0x4c, 0x43, 0x8a, 0x6e, 0x19, 0xf6, 0xff,
	0x23, 0x71, 0x5b, 0xe2, 0xb6, 0x62, 0x9f, 0xcc, 0x4f, 0xbb, 0xc2, 0x86, 0x69, 0xb0, 0x47, 0xf3,
	0x07, 0x4c, 0x84, 0xa0, 0x89, 0xa9, 0x2b, 0x2a, 0x89, 0x62, 0xfd, 0xb5, 0xab, 0x0e, 0xdf, 0x97,
	0xdf, 0xb7, 0x68, 0x17, 0x6e, 0x4f, 0xc6, 0x66, 0x54, 0x49, 0x50, 0x3a, 0x0d, 0xfd, 0x88, 0x1d,
	0x85, 0x8c, 0x75, 0xf5, 0x7a, 0xaf, 0xe6, 0xfa, 0x14, 0xe7, 0x90, 0xf1, 0x11, 0x00, 0x00, 0xff,
	0xff, 0xae, 0xe9, 0x01, 0x19, 0xcf, 0x00, 0x00, 0x00,
}
