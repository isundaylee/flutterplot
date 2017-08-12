// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flutter.proto

/*
Package flutter is a generated protocol buffer package.

It is generated from these files:
	flutter.proto

It has these top-level messages:
	DataPoint
	Data
*/
package flutter

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

type DataPoint struct {
	Entity string  `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
	Metric string  `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
	Value  float64 `protobuf:"fixed64,3,opt,name=value" json:"value,omitempty"`
}

func (m *DataPoint) Reset()                    { *m = DataPoint{} }
func (m *DataPoint) String() string            { return proto.CompactTextString(m) }
func (*DataPoint) ProtoMessage()               {}
func (*DataPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DataPoint) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *DataPoint) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *DataPoint) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Data struct {
	Points []*DataPoint `protobuf:"bytes,1,rep,name=points" json:"points,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Data) GetPoints() []*DataPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*DataPoint)(nil), "DataPoint")
	proto.RegisterType((*Data)(nil), "Data")
}

func init() { proto.RegisterFile("flutter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xcb, 0x29, 0x2d,
	0x29, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x0a, 0xe4, 0xe2, 0x74, 0x49, 0x2c,
	0x49, 0x0c, 0xc8, 0xcf, 0xcc, 0x2b, 0x11, 0x12, 0xe3, 0x62, 0x4b, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9,
	0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x40, 0xe2, 0xb9, 0xa9, 0x25, 0x45, 0x99,
	0xc9, 0x12, 0x4c, 0x10, 0x71, 0x08, 0x4f, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55,
	0x82, 0x59, 0x81, 0x51, 0x83, 0x31, 0x08, 0xc2, 0x51, 0xd2, 0xe2, 0x62, 0x01, 0x19, 0x29, 0xa4,
	0xc4, 0xc5, 0x56, 0x00, 0x32, 0xb6, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x88, 0x4b, 0x0f,
	0x6e, 0x53, 0x10, 0x54, 0x26, 0x89, 0x0d, 0xec, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x70, 0xef, 0x1b, 0x66, 0x96, 0x00, 0x00, 0x00,
}
