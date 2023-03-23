// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: bucketing/bucketing.proto

package bucketing

import (
	decision_response "github.com/flagship-io/flagship-proto/decision_response"
	targeting "github.com/flagship-io/flagship-proto/targeting"
	troubleshooting "github.com/flagship-io/flagship-proto/troubleshooting"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Bucketing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Bucketing) Reset() {
	*x = Bucketing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucketing_bucketing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucketing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucketing) ProtoMessage() {}

func (x *Bucketing) ProtoReflect() protoreflect.Message {
	mi := &file_bucketing_bucketing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucketing.ProtoReflect.Descriptor instead.
func (*Bucketing) Descriptor() ([]byte, []int) {
	return file_bucketing_bucketing_proto_rawDescGZIP(), []int{0}
}

type Bucketing_BucketingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GeneratedAt          string                             `protobuf:"bytes,1,opt,name=generated_at,json=generatedAt,proto3" json:"generated_at,omitempty"`
	VisitorConsolidation bool                               `protobuf:"varint,2,opt,name=visitor_consolidation,json=visitorConsolidation,proto3" json:"visitor_consolidation,omitempty"`
	Panic                bool                               `protobuf:"varint,3,opt,name=panic,proto3" json:"panic,omitempty"`
	Campaigns            []*Bucketing_BucketingCampaign     `protobuf:"bytes,4,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
	AccountSettings      *decision_response.AccountSettings `protobuf:"bytes,5,opt,name=accountSettings,proto3" json:"accountSettings,omitempty"`
	Troubleshooting      *troubleshooting.Troubleshooting   `protobuf:"bytes,6,opt,name=troubleshooting,proto3" json:"troubleshooting,omitempty"`
}

func (x *Bucketing_BucketingResponse) Reset() {
	*x = Bucketing_BucketingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucketing_bucketing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucketing_BucketingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucketing_BucketingResponse) ProtoMessage() {}

func (x *Bucketing_BucketingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucketing_bucketing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucketing_BucketingResponse.ProtoReflect.Descriptor instead.
func (*Bucketing_BucketingResponse) Descriptor() ([]byte, []int) {
	return file_bucketing_bucketing_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Bucketing_BucketingResponse) GetGeneratedAt() string {
	if x != nil {
		return x.GeneratedAt
	}
	return ""
}

func (x *Bucketing_BucketingResponse) GetVisitorConsolidation() bool {
	if x != nil {
		return x.VisitorConsolidation
	}
	return false
}

func (x *Bucketing_BucketingResponse) GetPanic() bool {
	if x != nil {
		return x.Panic
	}
	return false
}

func (x *Bucketing_BucketingResponse) GetCampaigns() []*Bucketing_BucketingCampaign {
	if x != nil {
		return x.Campaigns
	}
	return nil
}

func (x *Bucketing_BucketingResponse) GetAccountSettings() *decision_response.AccountSettings {
	if x != nil {
		return x.AccountSettings
	}
	return nil
}

func (x *Bucketing_BucketingResponse) GetTroubleshooting() *troubleshooting.Troubleshooting {
	if x != nil {
		return x.Troubleshooting
	}
	return nil
}

type Bucketing_BucketingCampaign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Deprecated: Do not use.
	CustomId        string                                     `protobuf:"bytes,2,opt,name=custom_id,json=customId,proto3" json:"custom_id,omitempty"`
	Slug            *wrapperspb.StringValue                    `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Type            string                                     `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	VariationGroups []*Bucketing_BucketingVariationGroups      `protobuf:"bytes,5,rep,name=variation_groups,json=variationGroups,proto3" json:"variation_groups,omitempty"`
	BucketRanges    []*Bucketing_BucketingCampaign_BucketRange `protobuf:"bytes,6,rep,name=bucket_ranges,json=bucketRanges,proto3" json:"bucket_ranges,omitempty"`
}

func (x *Bucketing_BucketingCampaign) Reset() {
	*x = Bucketing_BucketingCampaign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucketing_bucketing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucketing_BucketingCampaign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucketing_BucketingCampaign) ProtoMessage() {}

func (x *Bucketing_BucketingCampaign) ProtoReflect() protoreflect.Message {
	mi := &file_bucketing_bucketing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucketing_BucketingCampaign.ProtoReflect.Descriptor instead.
func (*Bucketing_BucketingCampaign) Descriptor() ([]byte, []int) {
	return file_bucketing_bucketing_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Bucketing_BucketingCampaign) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Deprecated: Do not use.
func (x *Bucketing_BucketingCampaign) GetCustomId() string {
	if x != nil {
		return x.CustomId
	}
	return ""
}

func (x *Bucketing_BucketingCampaign) GetSlug() *wrapperspb.StringValue {
	if x != nil {
		return x.Slug
	}
	return nil
}

func (x *Bucketing_BucketingCampaign) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Bucketing_BucketingCampaign) GetVariationGroups() []*Bucketing_BucketingVariationGroups {
	if x != nil {
		return x.VariationGroups
	}
	return nil
}

func (x *Bucketing_BucketingCampaign) GetBucketRanges() []*Bucketing_BucketingCampaign_BucketRange {
	if x != nil {
		return x.BucketRanges
	}
	return nil
}

type Bucketing_BucketingVariationGroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Targeting  *targeting.Targeting               `protobuf:"bytes,2,opt,name=targeting,proto3" json:"targeting,omitempty"`
	Variations []*decision_response.FullVariation `protobuf:"bytes,3,rep,name=variations,proto3" json:"variations,omitempty"`
}

func (x *Bucketing_BucketingVariationGroups) Reset() {
	*x = Bucketing_BucketingVariationGroups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucketing_bucketing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucketing_BucketingVariationGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucketing_BucketingVariationGroups) ProtoMessage() {}

func (x *Bucketing_BucketingVariationGroups) ProtoReflect() protoreflect.Message {
	mi := &file_bucketing_bucketing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucketing_BucketingVariationGroups.ProtoReflect.Descriptor instead.
func (*Bucketing_BucketingVariationGroups) Descriptor() ([]byte, []int) {
	return file_bucketing_bucketing_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Bucketing_BucketingVariationGroups) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bucketing_BucketingVariationGroups) GetTargeting() *targeting.Targeting {
	if x != nil {
		return x.Targeting
	}
	return nil
}

func (x *Bucketing_BucketingVariationGroups) GetVariations() []*decision_response.FullVariation {
	if x != nil {
		return x.Variations
	}
	return nil
}

type Bucketing_BucketingUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvironmentId string `protobuf:"bytes,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	RequestedAt   string `protobuf:"bytes,2,opt,name=requested_at,json=requestedAt,proto3" json:"requested_at,omitempty"`
}

func (x *Bucketing_BucketingUpdateRequest) Reset() {
	*x = Bucketing_BucketingUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucketing_bucketing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucketing_BucketingUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucketing_BucketingUpdateRequest) ProtoMessage() {}

func (x *Bucketing_BucketingUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bucketing_bucketing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucketing_BucketingUpdateRequest.ProtoReflect.Descriptor instead.
func (*Bucketing_BucketingUpdateRequest) Descriptor() ([]byte, []int) {
	return file_bucketing_bucketing_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Bucketing_BucketingUpdateRequest) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *Bucketing_BucketingUpdateRequest) GetRequestedAt() string {
	if x != nil {
		return x.RequestedAt
	}
	return ""
}

type Bucketing_BucketingCampaign_BucketRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R []float64 `protobuf:"fixed64,1,rep,packed,name=r,proto3" json:"r,omitempty"`
}

func (x *Bucketing_BucketingCampaign_BucketRange) Reset() {
	*x = Bucketing_BucketingCampaign_BucketRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucketing_bucketing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucketing_BucketingCampaign_BucketRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucketing_BucketingCampaign_BucketRange) ProtoMessage() {}

func (x *Bucketing_BucketingCampaign_BucketRange) ProtoReflect() protoreflect.Message {
	mi := &file_bucketing_bucketing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucketing_BucketingCampaign_BucketRange.ProtoReflect.Descriptor instead.
func (*Bucketing_BucketingCampaign_BucketRange) Descriptor() ([]byte, []int) {
	return file_bucketing_bucketing_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Bucketing_BucketingCampaign_BucketRange) GetR() []float64 {
	if x != nil {
		return x.R
	}
	return nil
}

var File_bucketing_bucketing_proto protoreflect.FileDescriptor

var file_bucketing_bucketing_proto_rawDesc = []byte{
	0x0a, 0x19, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x29,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x07, 0x0a, 0x09,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0xeb, 0x02, 0x0a, 0x11, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x33, 0x0a, 0x15, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x14, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x6e, 0x69, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x12, 0x4c, 0x0a,
	0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x52, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x74, 0x72, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73, 0x68,
	0x6f, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x74, 0x72, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x73,
	0x68, 0x6f, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0xea, 0x02, 0x0a, 0x11, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x30,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x60, 0x0a, 0x10, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x0f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x5f, 0x0a, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0x1b, 0x0a, 0x0b, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x01, 0x72, 0x1a, 0xa8, 0x01, 0x0a, 0x18, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a,
	0x0a, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x62, 0x0a, 0x16, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bucketing_bucketing_proto_rawDescOnce sync.Once
	file_bucketing_bucketing_proto_rawDescData = file_bucketing_bucketing_proto_rawDesc
)

func file_bucketing_bucketing_proto_rawDescGZIP() []byte {
	file_bucketing_bucketing_proto_rawDescOnce.Do(func() {
		file_bucketing_bucketing_proto_rawDescData = protoimpl.X.CompressGZIP(file_bucketing_bucketing_proto_rawDescData)
	})
	return file_bucketing_bucketing_proto_rawDescData
}

var file_bucketing_bucketing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bucketing_bucketing_proto_goTypes = []interface{}{
	(*Bucketing)(nil),                               // 0: flagship.protobuf.Bucketing
	(*Bucketing_BucketingResponse)(nil),             // 1: flagship.protobuf.Bucketing.BucketingResponse
	(*Bucketing_BucketingCampaign)(nil),             // 2: flagship.protobuf.Bucketing.BucketingCampaign
	(*Bucketing_BucketingVariationGroups)(nil),      // 3: flagship.protobuf.Bucketing.BucketingVariationGroups
	(*Bucketing_BucketingUpdateRequest)(nil),        // 4: flagship.protobuf.Bucketing.BucketingUpdateRequest
	(*Bucketing_BucketingCampaign_BucketRange)(nil), // 5: flagship.protobuf.Bucketing.BucketingCampaign.BucketRange
	(*decision_response.AccountSettings)(nil),       // 6: flagship.protobuf.AccountSettings
	(*troubleshooting.Troubleshooting)(nil),         // 7: flagship.protobuf.Troubleshooting
	(*wrapperspb.StringValue)(nil),                  // 8: google.protobuf.StringValue
	(*targeting.Targeting)(nil),                     // 9: flagship.protobuf.Targeting
	(*decision_response.FullVariation)(nil),         // 10: flagship.protobuf.FullVariation
}
var file_bucketing_bucketing_proto_depIdxs = []int32{
	2,  // 0: flagship.protobuf.Bucketing.BucketingResponse.campaigns:type_name -> flagship.protobuf.Bucketing.BucketingCampaign
	6,  // 1: flagship.protobuf.Bucketing.BucketingResponse.accountSettings:type_name -> flagship.protobuf.AccountSettings
	7,  // 2: flagship.protobuf.Bucketing.BucketingResponse.troubleshooting:type_name -> flagship.protobuf.Troubleshooting
	8,  // 3: flagship.protobuf.Bucketing.BucketingCampaign.slug:type_name -> google.protobuf.StringValue
	3,  // 4: flagship.protobuf.Bucketing.BucketingCampaign.variation_groups:type_name -> flagship.protobuf.Bucketing.BucketingVariationGroups
	5,  // 5: flagship.protobuf.Bucketing.BucketingCampaign.bucket_ranges:type_name -> flagship.protobuf.Bucketing.BucketingCampaign.BucketRange
	9,  // 6: flagship.protobuf.Bucketing.BucketingVariationGroups.targeting:type_name -> flagship.protobuf.Targeting
	10, // 7: flagship.protobuf.Bucketing.BucketingVariationGroups.variations:type_name -> flagship.protobuf.FullVariation
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_bucketing_bucketing_proto_init() }
func file_bucketing_bucketing_proto_init() {
	if File_bucketing_bucketing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bucketing_bucketing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucketing); i {
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
		file_bucketing_bucketing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucketing_BucketingResponse); i {
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
		file_bucketing_bucketing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucketing_BucketingCampaign); i {
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
		file_bucketing_bucketing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucketing_BucketingVariationGroups); i {
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
		file_bucketing_bucketing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucketing_BucketingUpdateRequest); i {
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
		file_bucketing_bucketing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucketing_BucketingCampaign_BucketRange); i {
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
			RawDescriptor: file_bucketing_bucketing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bucketing_bucketing_proto_goTypes,
		DependencyIndexes: file_bucketing_bucketing_proto_depIdxs,
		MessageInfos:      file_bucketing_bucketing_proto_msgTypes,
	}.Build()
	File_bucketing_bucketing_proto = out.File
	file_bucketing_bucketing_proto_rawDesc = nil
	file_bucketing_bucketing_proto_goTypes = nil
	file_bucketing_bucketing_proto_depIdxs = nil
}
