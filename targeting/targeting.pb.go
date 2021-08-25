// Code generated by protoc-gen-go. DO NOT EDIT.
// source: targeting/targeting.proto

package targeting

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Targeting_TargetingOperator int32

const (
	Targeting_NULL                   Targeting_TargetingOperator = 0
	Targeting_LOWER_THAN             Targeting_TargetingOperator = 1
	Targeting_GREATER_THAN_OR_EQUALS Targeting_TargetingOperator = 2
	Targeting_LOWER_THAN_OR_EQUALS   Targeting_TargetingOperator = 3
	Targeting_EQUALS                 Targeting_TargetingOperator = 4
	Targeting_NOT_EQUALS             Targeting_TargetingOperator = 5
	Targeting_STARTS_WITH            Targeting_TargetingOperator = 6
	Targeting_ENDS_WITH              Targeting_TargetingOperator = 7
	Targeting_CONTAINS               Targeting_TargetingOperator = 8
	Targeting_NOT_CONTAINS           Targeting_TargetingOperator = 9
	Targeting_GREATER_THAN           Targeting_TargetingOperator = 10
)

var Targeting_TargetingOperator_name = map[int32]string{
	0:  "NULL",
	1:  "LOWER_THAN",
	2:  "GREATER_THAN_OR_EQUALS",
	3:  "LOWER_THAN_OR_EQUALS",
	4:  "EQUALS",
	5:  "NOT_EQUALS",
	6:  "STARTS_WITH",
	7:  "ENDS_WITH",
	8:  "CONTAINS",
	9:  "NOT_CONTAINS",
	10: "GREATER_THAN",
}

var Targeting_TargetingOperator_value = map[string]int32{
	"NULL":                   0,
	"LOWER_THAN":             1,
	"GREATER_THAN_OR_EQUALS": 2,
	"LOWER_THAN_OR_EQUALS":   3,
	"EQUALS":                 4,
	"NOT_EQUALS":             5,
	"STARTS_WITH":            6,
	"ENDS_WITH":              7,
	"CONTAINS":               8,
	"NOT_CONTAINS":           9,
	"GREATER_THAN":           10,
}

func (x Targeting_TargetingOperator) String() string {
	return proto.EnumName(Targeting_TargetingOperator_name, int32(x))
}

func (Targeting_TargetingOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89898cd1f7050311, []int{0, 0}
}

type Targeting struct {
	TargetingGroups      []*Targeting_TargetingGroup `protobuf:"bytes,1,rep,name=targeting_groups,json=targetingGroups,proto3" json:"targeting_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Targeting) Reset()         { *m = Targeting{} }
func (m *Targeting) String() string { return proto.CompactTextString(m) }
func (*Targeting) ProtoMessage()    {}
func (*Targeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_89898cd1f7050311, []int{0}
}

func (m *Targeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Targeting.Unmarshal(m, b)
}
func (m *Targeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Targeting.Marshal(b, m, deterministic)
}
func (m *Targeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Targeting.Merge(m, src)
}
func (m *Targeting) XXX_Size() int {
	return xxx_messageInfo_Targeting.Size(m)
}
func (m *Targeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Targeting.DiscardUnknown(m)
}

var xxx_messageInfo_Targeting proto.InternalMessageInfo

func (m *Targeting) GetTargetingGroups() []*Targeting_TargetingGroup {
	if m != nil {
		return m.TargetingGroups
	}
	return nil
}

type Targeting_InnerTargeting struct {
	//*
	// targeting operator
	//
	// Required
	//
	// The operator for the targeting
	Operator Targeting_TargetingOperator `protobuf:"varint,1,opt,name=operator,proto3,enum=canarybay.protobuf.Targeting_TargetingOperator" json:"operator,omitempty"`
	//*
	// key
	//
	// Required
	//
	// The targeting key for the context
	Key *wrappers.StringValue `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	//*
	// value
	//
	// Required
	//
	// The targeting value
	Value                *_struct.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Targeting_InnerTargeting) Reset()         { *m = Targeting_InnerTargeting{} }
func (m *Targeting_InnerTargeting) String() string { return proto.CompactTextString(m) }
func (*Targeting_InnerTargeting) ProtoMessage()    {}
func (*Targeting_InnerTargeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_89898cd1f7050311, []int{0, 0}
}

func (m *Targeting_InnerTargeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Targeting_InnerTargeting.Unmarshal(m, b)
}
func (m *Targeting_InnerTargeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Targeting_InnerTargeting.Marshal(b, m, deterministic)
}
func (m *Targeting_InnerTargeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Targeting_InnerTargeting.Merge(m, src)
}
func (m *Targeting_InnerTargeting) XXX_Size() int {
	return xxx_messageInfo_Targeting_InnerTargeting.Size(m)
}
func (m *Targeting_InnerTargeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Targeting_InnerTargeting.DiscardUnknown(m)
}

var xxx_messageInfo_Targeting_InnerTargeting proto.InternalMessageInfo

func (m *Targeting_InnerTargeting) GetOperator() Targeting_TargetingOperator {
	if m != nil {
		return m.Operator
	}
	return Targeting_NULL
}

func (m *Targeting_InnerTargeting) GetKey() *wrappers.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Targeting_InnerTargeting) GetValue() *_struct.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Targeting_TargetingGroup struct {
	//*
	// visitor ID
	//
	// Required
	//
	// The custom or decision generated visitor ID.
	Targetings           []*Targeting_InnerTargeting `protobuf:"bytes,1,rep,name=targetings,proto3" json:"targetings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Targeting_TargetingGroup) Reset()         { *m = Targeting_TargetingGroup{} }
func (m *Targeting_TargetingGroup) String() string { return proto.CompactTextString(m) }
func (*Targeting_TargetingGroup) ProtoMessage()    {}
func (*Targeting_TargetingGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_89898cd1f7050311, []int{0, 1}
}

func (m *Targeting_TargetingGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Targeting_TargetingGroup.Unmarshal(m, b)
}
func (m *Targeting_TargetingGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Targeting_TargetingGroup.Marshal(b, m, deterministic)
}
func (m *Targeting_TargetingGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Targeting_TargetingGroup.Merge(m, src)
}
func (m *Targeting_TargetingGroup) XXX_Size() int {
	return xxx_messageInfo_Targeting_TargetingGroup.Size(m)
}
func (m *Targeting_TargetingGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_Targeting_TargetingGroup.DiscardUnknown(m)
}

var xxx_messageInfo_Targeting_TargetingGroup proto.InternalMessageInfo

func (m *Targeting_TargetingGroup) GetTargetings() []*Targeting_InnerTargeting {
	if m != nil {
		return m.Targetings
	}
	return nil
}

func init() {
	proto.RegisterEnum("canarybay.protobuf.Targeting_TargetingOperator", Targeting_TargetingOperator_name, Targeting_TargetingOperator_value)
	proto.RegisterType((*Targeting)(nil), "canarybay.protobuf.Targeting")
	proto.RegisterType((*Targeting_InnerTargeting)(nil), "canarybay.protobuf.Targeting.InnerTargeting")
	proto.RegisterType((*Targeting_TargetingGroup)(nil), "canarybay.protobuf.Targeting.TargetingGroup")
}

func init() { proto.RegisterFile("targeting/targeting.proto", fileDescriptor_89898cd1f7050311) }

var fileDescriptor_89898cd1f7050311 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x75, 0x80, 0x45, 0xb8, 0xac, 0xec, 0x38, 0x31, 0xa4, 0x36, 0x1b, 0x43, 0xf6, 0x89, 0x87,
	0x4d, 0x9b, 0xe0, 0x17, 0x94, 0x95, 0xb0, 0xc4, 0xda, 0x62, 0x3b, 0x2b, 0x89, 0xd9, 0xd8, 0x0c,
	0x38, 0x36, 0x44, 0xb6, 0xd3, 0x4c, 0x07, 0x4d, 0x7f, 0xc2, 0x1f, 0xf2, 0x0f, 0x7c, 0xf3, 0xc5,
	0x0f, 0xf0, 0x4b, 0x4c, 0x69, 0x69, 0x61, 0xf1, 0x81, 0xb7, 0x7b, 0xcf, 0x3d, 0xe7, 0x9e, 0x9e,
	0x9b, 0x0e, 0xbc, 0x54, 0x4c, 0x86, 0x5c, 0xad, 0xa2, 0xd0, 0x2c, 0x2b, 0x23, 0x96, 0x42, 0x09,
	0x42, 0x96, 0x2c, 0x62, 0x32, 0x5d, 0xb0, 0x34, 0x07, 0x16, 0x9b, 0x2f, 0xfa, 0xab, 0x50, 0x88,
	0x70, 0xcd, 0xcd, 0x1d, 0x60, 0x7e, 0x97, 0x2c, 0x8e, 0xb9, 0x4c, 0x72, 0x8a, 0x7e, 0xf9, 0x78,
	0x9e, 0x28, 0xb9, 0x59, 0xaa, 0x7c, 0x7a, 0xf5, 0xa7, 0x01, 0x6d, 0xba, 0x73, 0x21, 0x73, 0xc0,
	0xa5, 0x65, 0x10, 0x4a, 0xb1, 0x89, 0x13, 0x0d, 0xf5, 0xeb, 0x83, 0xce, 0xf0, 0xda, 0x38, 0xb6,
	0x36, 0x4a, 0x61, 0x55, 0x4d, 0x32, 0x91, 0x77, 0xa1, 0x0e, 0xfa, 0x44, 0xff, 0x89, 0xa0, 0x3b,
	0x8d, 0x22, 0x2e, 0x2b, 0xaf, 0xb7, 0xd0, 0x12, 0x31, 0x97, 0x4c, 0x09, 0xa9, 0xa1, 0x3e, 0x1a,
	0x74, 0x87, 0xe6, 0x89, 0x1e, 0x6e, 0x21, 0xf3, 0xca, 0x05, 0xc4, 0x80, 0xfa, 0x57, 0x9e, 0x6a,
	0xb5, 0x3e, 0x1a, 0x74, 0x86, 0x97, 0x46, 0x1e, 0xb9, 0x5a, 0xe2, 0x2b, 0xb9, 0x8a, 0xc2, 0x0f,
	0x6c, 0xbd, 0xe1, 0x5e, 0x46, 0x24, 0xd7, 0x70, 0xf6, 0x2d, 0xeb, 0xb4, 0xfa, 0x56, 0xd1, 0x3b,
	0x52, 0xe4, 0xdc, 0x9c, 0xa4, 0x7f, 0x82, 0xee, 0x61, 0x40, 0x62, 0x03, 0x94, 0x11, 0x4f, 0x3c,
	0xd1, 0x61, 0x7c, 0x6f, 0x4f, 0x7f, 0xf5, 0x1b, 0xc1, 0xf3, 0xa3, 0x74, 0xa4, 0x05, 0x0d, 0xe7,
	0xce, 0xb6, 0xf1, 0x13, 0xd2, 0x05, 0xb0, 0xdd, 0xf9, 0xd8, 0x0b, 0xe8, 0xad, 0xe5, 0x60, 0x44,
	0x74, 0xe8, 0x4d, 0xbc, 0xb1, 0x45, 0x0b, 0x24, 0x70, 0xbd, 0x60, 0xfc, 0xfe, 0xce, 0xb2, 0x7d,
	0x5c, 0x23, 0x1a, 0xbc, 0xa8, 0xb8, 0x7b, 0x93, 0x3a, 0x01, 0x68, 0x16, 0x75, 0x23, 0xdb, 0xe8,
	0xb8, 0x74, 0x37, 0x3b, 0x23, 0x17, 0xd0, 0xf1, 0xa9, 0xe5, 0x51, 0x3f, 0x98, 0x4f, 0xe9, 0x2d,
	0x6e, 0x92, 0x67, 0xd0, 0x1e, 0x3b, 0x6f, 0x8a, 0xf6, 0x29, 0x39, 0x87, 0xd6, 0x8d, 0xeb, 0x50,
	0x6b, 0xea, 0xf8, 0xb8, 0x45, 0x30, 0x9c, 0x67, 0xea, 0x12, 0x69, 0x67, 0xc8, 0xfe, 0x17, 0x61,
	0x18, 0xfd, 0x40, 0xd0, 0x5b, 0x8a, 0x87, 0xff, 0xdc, 0x64, 0x54, 0x1d, 0x73, 0x96, 0x41, 0x33,
	0xf4, 0x71, 0x18, 0xae, 0xd4, 0x9a, 0x2d, 0x8c, 0xa5, 0x78, 0x30, 0x4b, 0x41, 0xf5, 0xc7, 0xc6,
	0x2a, 0x8d, 0x79, 0x62, 0x84, 0x2b, 0x55, 0xbd, 0x88, 0x5f, 0x35, 0x7c, 0xb3, 0x65, 0x8e, 0x58,
	0x7a, 0x3f, 0xa3, 0x19, 0xe1, 0x6f, 0x4d, 0x7f, 0x0c, 0xdd, 0x4f, 0x66, 0xa3, 0x77, 0x5c, 0xb1,
	0xcf, 0x6c, 0xd1, 0xdc, 0xae, 0x7b, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x82, 0xc8, 0xcc, 0xc1,
	0x5f, 0x03, 0x00, 0x00,
}
