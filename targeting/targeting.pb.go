// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: targeting/targeting.proto

package targeting

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

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
	Targeting_EXISTS                 Targeting_TargetingOperator = 11
	Targeting_NOT_EXISTS             Targeting_TargetingOperator = 12
)

// Enum value maps for Targeting_TargetingOperator.
var (
	Targeting_TargetingOperator_name = map[int32]string{
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
		11: "EXISTS",
		12: "NOT_EXISTS",
	}
	Targeting_TargetingOperator_value = map[string]int32{
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
		"EXISTS":                 11,
		"NOT_EXISTS":             12,
	}
)

func (x Targeting_TargetingOperator) Enum() *Targeting_TargetingOperator {
	p := new(Targeting_TargetingOperator)
	*p = x
	return p
}

func (x Targeting_TargetingOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Targeting_TargetingOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_targeting_targeting_proto_enumTypes[0].Descriptor()
}

func (Targeting_TargetingOperator) Type() protoreflect.EnumType {
	return &file_targeting_targeting_proto_enumTypes[0]
}

func (x Targeting_TargetingOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Targeting_TargetingOperator.Descriptor instead.
func (Targeting_TargetingOperator) EnumDescriptor() ([]byte, []int) {
	return file_targeting_targeting_proto_rawDescGZIP(), []int{0, 0}
}

type Targeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetingGroups []*Targeting_TargetingGroup `protobuf:"bytes,1,rep,name=targeting_groups,json=targetingGroups,proto3" json:"targeting_groups,omitempty"`
}

func (x *Targeting) Reset() {
	*x = Targeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_targeting_targeting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Targeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Targeting) ProtoMessage() {}

func (x *Targeting) ProtoReflect() protoreflect.Message {
	mi := &file_targeting_targeting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Targeting.ProtoReflect.Descriptor instead.
func (*Targeting) Descriptor() ([]byte, []int) {
	return file_targeting_targeting_proto_rawDescGZIP(), []int{0}
}

func (x *Targeting) GetTargetingGroups() []*Targeting_TargetingGroup {
	if x != nil {
		return x.TargetingGroups
	}
	return nil
}

type Targeting_InnerTargeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// targeting operator
	//
	// # Required
	//
	// The operator for the targeting
	Operator Targeting_TargetingOperator `protobuf:"varint,1,opt,name=operator,proto3,enum=flagship.protobuf.Targeting_TargetingOperator" json:"operator,omitempty"`
	// *
	// key
	//
	// # Required
	//
	// The targeting key for the context
	Key *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// *
	// value
	//
	// # Required
	//
	// The targeting value
	Value *structpb.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// *
	// provider
	//
	// # Optional
	//
	// The integration provider key
	Provider *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *Targeting_InnerTargeting) Reset() {
	*x = Targeting_InnerTargeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_targeting_targeting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Targeting_InnerTargeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Targeting_InnerTargeting) ProtoMessage() {}

func (x *Targeting_InnerTargeting) ProtoReflect() protoreflect.Message {
	mi := &file_targeting_targeting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Targeting_InnerTargeting.ProtoReflect.Descriptor instead.
func (*Targeting_InnerTargeting) Descriptor() ([]byte, []int) {
	return file_targeting_targeting_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Targeting_InnerTargeting) GetOperator() Targeting_TargetingOperator {
	if x != nil {
		return x.Operator
	}
	return Targeting_NULL
}

func (x *Targeting_InnerTargeting) GetKey() *wrapperspb.StringValue {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Targeting_InnerTargeting) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Targeting_InnerTargeting) GetProvider() *wrapperspb.StringValue {
	if x != nil {
		return x.Provider
	}
	return nil
}

type Targeting_TargetingGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// visitor ID
	//
	// # Required
	//
	// The custom or decision generated visitor ID.
	Targetings []*Targeting_InnerTargeting `protobuf:"bytes,1,rep,name=targetings,proto3" json:"targetings,omitempty"`
}

func (x *Targeting_TargetingGroup) Reset() {
	*x = Targeting_TargetingGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_targeting_targeting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Targeting_TargetingGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Targeting_TargetingGroup) ProtoMessage() {}

func (x *Targeting_TargetingGroup) ProtoReflect() protoreflect.Message {
	mi := &file_targeting_targeting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Targeting_TargetingGroup.ProtoReflect.Descriptor instead.
func (*Targeting_TargetingGroup) Descriptor() ([]byte, []int) {
	return file_targeting_targeting_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Targeting_TargetingGroup) GetTargetings() []*Targeting_InnerTargeting {
	if x != nil {
		return x.Targetings
	}
	return nil
}

var File_targeting_targeting_proto protoreflect.FileDescriptor

var file_targeting_targeting_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x05, 0x0a,
	0x09, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x56, 0x0a, 0x10, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x1a, 0xf4, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x4a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68,
	0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x2e, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x5d, 0x0a, 0x0e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x4b, 0x0a, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xed, 0x01, 0x0a, 0x11, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x57, 0x45,
	0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x55, 0x41,
	0x4c, 0x53, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x54, 0x48,
	0x41, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f,
	0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x45,
	0x4e, 0x44, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f,
	0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x0c, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2d,
	0x69, 0x6f, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_targeting_targeting_proto_rawDescOnce sync.Once
	file_targeting_targeting_proto_rawDescData = file_targeting_targeting_proto_rawDesc
)

func file_targeting_targeting_proto_rawDescGZIP() []byte {
	file_targeting_targeting_proto_rawDescOnce.Do(func() {
		file_targeting_targeting_proto_rawDescData = protoimpl.X.CompressGZIP(file_targeting_targeting_proto_rawDescData)
	})
	return file_targeting_targeting_proto_rawDescData
}

var file_targeting_targeting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_targeting_targeting_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_targeting_targeting_proto_goTypes = []interface{}{
	(Targeting_TargetingOperator)(0), // 0: flagship.protobuf.Targeting.TargetingOperator
	(*Targeting)(nil),                // 1: flagship.protobuf.Targeting
	(*Targeting_InnerTargeting)(nil), // 2: flagship.protobuf.Targeting.InnerTargeting
	(*Targeting_TargetingGroup)(nil), // 3: flagship.protobuf.Targeting.TargetingGroup
	(*wrapperspb.StringValue)(nil),   // 4: google.protobuf.StringValue
	(*structpb.Value)(nil),           // 5: google.protobuf.Value
}
var file_targeting_targeting_proto_depIdxs = []int32{
	3, // 0: flagship.protobuf.Targeting.targeting_groups:type_name -> flagship.protobuf.Targeting.TargetingGroup
	0, // 1: flagship.protobuf.Targeting.InnerTargeting.operator:type_name -> flagship.protobuf.Targeting.TargetingOperator
	4, // 2: flagship.protobuf.Targeting.InnerTargeting.key:type_name -> google.protobuf.StringValue
	5, // 3: flagship.protobuf.Targeting.InnerTargeting.value:type_name -> google.protobuf.Value
	4, // 4: flagship.protobuf.Targeting.InnerTargeting.provider:type_name -> google.protobuf.StringValue
	2, // 5: flagship.protobuf.Targeting.TargetingGroup.targetings:type_name -> flagship.protobuf.Targeting.InnerTargeting
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_targeting_targeting_proto_init() }
func file_targeting_targeting_proto_init() {
	if File_targeting_targeting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_targeting_targeting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Targeting); i {
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
		file_targeting_targeting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Targeting_InnerTargeting); i {
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
		file_targeting_targeting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Targeting_TargetingGroup); i {
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
			RawDescriptor: file_targeting_targeting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_targeting_targeting_proto_goTypes,
		DependencyIndexes: file_targeting_targeting_proto_depIdxs,
		EnumInfos:         file_targeting_targeting_proto_enumTypes,
		MessageInfos:      file_targeting_targeting_proto_msgTypes,
	}.Build()
	File_targeting_targeting_proto = out.File
	file_targeting_targeting_proto_rawDesc = nil
	file_targeting_targeting_proto_goTypes = nil
	file_targeting_targeting_proto_depIdxs = nil
}
