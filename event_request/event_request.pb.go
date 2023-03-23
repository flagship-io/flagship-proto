// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: event_request/event_request.proto

package event_request

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

type EventRequest_EventType int32

const (
	EventRequest_NULL    EventRequest_EventType = 0
	EventRequest_CONTEXT EventRequest_EventType = 1
)

// Enum value maps for EventRequest_EventType.
var (
	EventRequest_EventType_name = map[int32]string{
		0: "NULL",
		1: "CONTEXT",
	}
	EventRequest_EventType_value = map[string]int32{
		"NULL":    0,
		"CONTEXT": 1,
	}
)

func (x EventRequest_EventType) Enum() *EventRequest_EventType {
	p := new(EventRequest_EventType)
	*p = x
	return p
}

func (x EventRequest_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventRequest_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_request_event_request_proto_enumTypes[0].Descriptor()
}

func (EventRequest_EventType) Type() protoreflect.EnumType {
	return &file_event_request_event_request_proto_enumTypes[0]
}

func (x EventRequest_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventRequest_EventType.Descriptor instead.
func (EventRequest_EventType) EnumDescriptor() ([]byte, []int) {
	return file_event_request_event_request_proto_rawDescGZIP(), []int{0, 0}
}

type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// client ID
	//
	// # Required
	//
	// The client ID.
	ClientId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// *
	// visitor ID
	//
	// # Optional
	//
	// The custom visitor ID.
	VisitorId *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=visitor_id,json=visitorId,proto3" json:"visitor_id,omitempty"`
	// *
	// type
	//
	// required
	//
	// The event type
	Type EventRequest_EventType `protobuf:"varint,3,opt,name=type,proto3,enum=flagship.protobuf.EventRequest_EventType" json:"type,omitempty"`
	// *
	// data
	//
	// # Required
	//
	// The event data for the current event request
	Data map[string]*structpb.Value `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_request_event_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_request_event_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_event_request_event_request_proto_rawDescGZIP(), []int{0}
}

func (x *EventRequest) GetClientId() *wrapperspb.StringValue {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *EventRequest) GetVisitorId() *wrapperspb.StringValue {
	if x != nil {
		return x.VisitorId
	}
	return nil
}

func (x *EventRequest) GetType() EventRequest_EventType {
	if x != nil {
		return x.Type
	}
	return EventRequest_NULL
}

func (x *EventRequest) GetData() map[string]*structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_event_request_event_request_proto protoreflect.FileDescriptor

var file_event_request_event_request_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x02, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x3b, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x4f, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c,
	0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x68, 0x69, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_request_event_request_proto_rawDescOnce sync.Once
	file_event_request_event_request_proto_rawDescData = file_event_request_event_request_proto_rawDesc
)

func file_event_request_event_request_proto_rawDescGZIP() []byte {
	file_event_request_event_request_proto_rawDescOnce.Do(func() {
		file_event_request_event_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_request_event_request_proto_rawDescData)
	})
	return file_event_request_event_request_proto_rawDescData
}

var file_event_request_event_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_event_request_event_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_event_request_event_request_proto_goTypes = []interface{}{
	(EventRequest_EventType)(0),    // 0: flagship.protobuf.EventRequest.EventType
	(*EventRequest)(nil),           // 1: flagship.protobuf.EventRequest
	nil,                            // 2: flagship.protobuf.EventRequest.DataEntry
	(*wrapperspb.StringValue)(nil), // 3: google.protobuf.StringValue
	(*structpb.Value)(nil),         // 4: google.protobuf.Value
}
var file_event_request_event_request_proto_depIdxs = []int32{
	3, // 0: flagship.protobuf.EventRequest.client_id:type_name -> google.protobuf.StringValue
	3, // 1: flagship.protobuf.EventRequest.visitor_id:type_name -> google.protobuf.StringValue
	0, // 2: flagship.protobuf.EventRequest.type:type_name -> flagship.protobuf.EventRequest.EventType
	2, // 3: flagship.protobuf.EventRequest.data:type_name -> flagship.protobuf.EventRequest.DataEntry
	4, // 4: flagship.protobuf.EventRequest.DataEntry.value:type_name -> google.protobuf.Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_event_request_event_request_proto_init() }
func file_event_request_event_request_proto_init() {
	if File_event_request_event_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_request_event_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRequest); i {
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
			RawDescriptor: file_event_request_event_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_request_event_request_proto_goTypes,
		DependencyIndexes: file_event_request_event_request_proto_depIdxs,
		EnumInfos:         file_event_request_event_request_proto_enumTypes,
		MessageInfos:      file_event_request_event_request_proto_msgTypes,
	}.Build()
	File_event_request_event_request_proto = out.File
	file_event_request_event_request_proto_rawDesc = nil
	file_event_request_event_request_proto_goTypes = nil
	file_event_request_event_request_proto_depIdxs = nil
}
