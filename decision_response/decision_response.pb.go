// Code generated by protoc-gen-go. DO NOT EDIT.
// source: decision_response/decision_response.proto

package decision_response

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

type ModificationsType int32

const (
	ModificationsType_NULL          ModificationsType = 0
	ModificationsType_JSON          ModificationsType = 1
	ModificationsType_TEXT          ModificationsType = 2
	ModificationsType_IMAGE         ModificationsType = 3
	ModificationsType_HTML          ModificationsType = 4
	ModificationsType_FLAG          ModificationsType = 5
	ModificationsType_REDIRECT      ModificationsType = 6
	ModificationsType_TURING_ENGINE ModificationsType = 7
)

var ModificationsType_name = map[int32]string{
	0: "NULL",
	1: "JSON",
	2: "TEXT",
	3: "IMAGE",
	4: "HTML",
	5: "FLAG",
	6: "REDIRECT",
	7: "TURING_ENGINE",
}

var ModificationsType_value = map[string]int32{
	"NULL":          0,
	"JSON":          1,
	"TEXT":          2,
	"IMAGE":         3,
	"HTML":          4,
	"FLAG":          5,
	"REDIRECT":      6,
	"TURING_ENGINE": 7,
}

func (x ModificationsType) String() string {
	return proto.EnumName(ModificationsType_name, int32(x))
}

func (ModificationsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{0}
}

type CampaignIdVariationId struct {
	//*
	// campaign id
	//
	// Required
	//
	// The campaign id
	CampaignId string `protobuf:"bytes,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	//*
	// variation id
	//
	// Required
	//
	// The variation id
	VariationId          string   `protobuf:"bytes,2,opt,name=variation_id,json=variationId,proto3" json:"variation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignIdVariationId) Reset()         { *m = CampaignIdVariationId{} }
func (m *CampaignIdVariationId) String() string { return proto.CompactTextString(m) }
func (*CampaignIdVariationId) ProtoMessage()    {}
func (*CampaignIdVariationId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{0}
}

func (m *CampaignIdVariationId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignIdVariationId.Unmarshal(m, b)
}
func (m *CampaignIdVariationId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignIdVariationId.Marshal(b, m, deterministic)
}
func (m *CampaignIdVariationId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignIdVariationId.Merge(m, src)
}
func (m *CampaignIdVariationId) XXX_Size() int {
	return xxx_messageInfo_CampaignIdVariationId.Size(m)
}
func (m *CampaignIdVariationId) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignIdVariationId.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignIdVariationId proto.InternalMessageInfo

func (m *CampaignIdVariationId) GetCampaignId() string {
	if m != nil {
		return m.CampaignId
	}
	return ""
}

func (m *CampaignIdVariationId) GetVariationId() string {
	if m != nil {
		return m.VariationId
	}
	return ""
}

type Modifications struct {
	//*
	// modification type
	//
	// Required
	//
	// The type of the modification of the variation
	Type ModificationsType `protobuf:"varint,1,opt,name=type,proto3,enum=flagship.protobuf.ModificationsType" json:"type,omitempty"`
	//*
	// value
	//
	// Required
	//
	// The modifications value in json
	Value                *_struct.Struct `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Modifications) Reset()         { *m = Modifications{} }
func (m *Modifications) String() string { return proto.CompactTextString(m) }
func (*Modifications) ProtoMessage()    {}
func (*Modifications) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{1}
}

func (m *Modifications) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Modifications.Unmarshal(m, b)
}
func (m *Modifications) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Modifications.Marshal(b, m, deterministic)
}
func (m *Modifications) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Modifications.Merge(m, src)
}
func (m *Modifications) XXX_Size() int {
	return xxx_messageInfo_Modifications.Size(m)
}
func (m *Modifications) XXX_DiscardUnknown() {
	xxx_messageInfo_Modifications.DiscardUnknown(m)
}

var xxx_messageInfo_Modifications proto.InternalMessageInfo

func (m *Modifications) GetType() ModificationsType {
	if m != nil {
		return m.Type
	}
	return ModificationsType_NULL
}

func (m *Modifications) GetValue() *_struct.Struct {
	if m != nil {
		return m.Value
	}
	return nil
}

type TuringEngineOptions struct {
	//*
	// Turing engine Context
	//
	// Required
	//
	// The context key(s) sent to turing engine
	ContextKeys          map[string]string `protobuf:"bytes,1,rep,name=context_keys,json=contextKeys,proto3" json:"context_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TuringEngineOptions) Reset()         { *m = TuringEngineOptions{} }
func (m *TuringEngineOptions) String() string { return proto.CompactTextString(m) }
func (*TuringEngineOptions) ProtoMessage()    {}
func (*TuringEngineOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{2}
}

func (m *TuringEngineOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TuringEngineOptions.Unmarshal(m, b)
}
func (m *TuringEngineOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TuringEngineOptions.Marshal(b, m, deterministic)
}
func (m *TuringEngineOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TuringEngineOptions.Merge(m, src)
}
func (m *TuringEngineOptions) XXX_Size() int {
	return xxx_messageInfo_TuringEngineOptions.Size(m)
}
func (m *TuringEngineOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TuringEngineOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TuringEngineOptions proto.InternalMessageInfo

func (m *TuringEngineOptions) GetContextKeys() map[string]string {
	if m != nil {
		return m.ContextKeys
	}
	return nil
}

type Variation struct {
	//*
	// variation ID
	//
	// Required
	//
	// The ID of the variation the visitor is affected to
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//*
	// modifications
	//
	// Required
	//
	// The modifications field for the variation
	Modifications *Modifications `protobuf:"bytes,2,opt,name=modifications,proto3" json:"modifications,omitempty"`
	//*
	// is reference
	//
	// Optional
	//
	// True if variation is reference
	Reference            bool     `protobuf:"varint,3,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Variation) Reset()         { *m = Variation{} }
func (m *Variation) String() string { return proto.CompactTextString(m) }
func (*Variation) ProtoMessage()    {}
func (*Variation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{3}
}

func (m *Variation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Variation.Unmarshal(m, b)
}
func (m *Variation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Variation.Marshal(b, m, deterministic)
}
func (m *Variation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Variation.Merge(m, src)
}
func (m *Variation) XXX_Size() int {
	return xxx_messageInfo_Variation.Size(m)
}
func (m *Variation) XXX_DiscardUnknown() {
	xxx_messageInfo_Variation.DiscardUnknown(m)
}

var xxx_messageInfo_Variation proto.InternalMessageInfo

func (m *Variation) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Variation) GetModifications() *Modifications {
	if m != nil {
		return m.Modifications
	}
	return nil
}

func (m *Variation) GetReference() bool {
	if m != nil {
		return m.Reference
	}
	return false
}

type FullVariation struct {
	//*
	// variation ID
	//
	// Required
	//
	// The ID of the variation the visitor is affected to
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//*
	// modifications
	//
	// Required
	//
	// The modifications field for the variation
	Modifications *Modifications `protobuf:"bytes,2,opt,name=modifications,proto3" json:"modifications,omitempty"`
	//*
	// Turing engine Options
	//
	// Optional
	//
	// The turing engine options
	TeOptions *TuringEngineOptions `protobuf:"bytes,3,opt,name=te_options,json=teOptions,proto3" json:"te_options,omitempty"`
	//*
	// Allocation
	//
	// Optional
	//
	// The allocation
	Allocation int32 `protobuf:"varint,4,opt,name=allocation,proto3" json:"allocation,omitempty"`
	//*
	// is reference
	//
	// Optional
	//
	// True if variation is reference
	Reference            bool     `protobuf:"varint,5,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FullVariation) Reset()         { *m = FullVariation{} }
func (m *FullVariation) String() string { return proto.CompactTextString(m) }
func (*FullVariation) ProtoMessage()    {}
func (*FullVariation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{4}
}

func (m *FullVariation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FullVariation.Unmarshal(m, b)
}
func (m *FullVariation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FullVariation.Marshal(b, m, deterministic)
}
func (m *FullVariation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FullVariation.Merge(m, src)
}
func (m *FullVariation) XXX_Size() int {
	return xxx_messageInfo_FullVariation.Size(m)
}
func (m *FullVariation) XXX_DiscardUnknown() {
	xxx_messageInfo_FullVariation.DiscardUnknown(m)
}

var xxx_messageInfo_FullVariation proto.InternalMessageInfo

func (m *FullVariation) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FullVariation) GetModifications() *Modifications {
	if m != nil {
		return m.Modifications
	}
	return nil
}

func (m *FullVariation) GetTeOptions() *TuringEngineOptions {
	if m != nil {
		return m.TeOptions
	}
	return nil
}

func (m *FullVariation) GetAllocation() int32 {
	if m != nil {
		return m.Allocation
	}
	return 0
}

func (m *FullVariation) GetReference() bool {
	if m != nil {
		return m.Reference
	}
	return false
}

type Campaign struct {
	//*
	// campaign ID
	//
	// Required
	//
	// The ID of the campaign the visitor is affected to
	Id *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//*
	// variation group ID
	//
	// Required
	//
	// The ID of the variation group the visitor is affected to
	VariationGroupId *wrappers.StringValue `protobuf:"bytes,2,opt,name=variation_group_id,json=variationGroupId,proto3" json:"variation_group_id,omitempty"`
	//*
	// variation
	//
	// Required
	//
	// The variation info the visitor is affected to
	Variation            *Variation `protobuf:"bytes,3,opt,name=variation,proto3" json:"variation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Campaign) Reset()         { *m = Campaign{} }
func (m *Campaign) String() string { return proto.CompactTextString(m) }
func (*Campaign) ProtoMessage()    {}
func (*Campaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{5}
}

func (m *Campaign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Campaign.Unmarshal(m, b)
}
func (m *Campaign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Campaign.Marshal(b, m, deterministic)
}
func (m *Campaign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Campaign.Merge(m, src)
}
func (m *Campaign) XXX_Size() int {
	return xxx_messageInfo_Campaign.Size(m)
}
func (m *Campaign) XXX_DiscardUnknown() {
	xxx_messageInfo_Campaign.DiscardUnknown(m)
}

var xxx_messageInfo_Campaign proto.InternalMessageInfo

func (m *Campaign) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Campaign) GetVariationGroupId() *wrappers.StringValue {
	if m != nil {
		return m.VariationGroupId
	}
	return nil
}

func (m *Campaign) GetVariation() *Variation {
	if m != nil {
		return m.Variation
	}
	return nil
}

type AccountSettings struct {
	//*
	// Experience continuity
	//
	// Required
	//
	// True if experience continuity is activated in account settings
	EnabledXPX bool `protobuf:"varint,1,opt,name=enabledXPX,proto3" json:"enabledXPX,omitempty"`
	//*
	// One visitor one test
	//
	// Required
	//
	// True if one visitor one test is activated in account settings
	Enabled1V1T          bool     `protobuf:"varint,2,opt,name=enabled1V1T,proto3" json:"enabled1V1T,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountSettings) Reset()         { *m = AccountSettings{} }
func (m *AccountSettings) String() string { return proto.CompactTextString(m) }
func (*AccountSettings) ProtoMessage()    {}
func (*AccountSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{6}
}

func (m *AccountSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSettings.Unmarshal(m, b)
}
func (m *AccountSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSettings.Marshal(b, m, deterministic)
}
func (m *AccountSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSettings.Merge(m, src)
}
func (m *AccountSettings) XXX_Size() int {
	return xxx_messageInfo_AccountSettings.Size(m)
}
func (m *AccountSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSettings proto.InternalMessageInfo

func (m *AccountSettings) GetEnabledXPX() bool {
	if m != nil {
		return m.EnabledXPX
	}
	return false
}

func (m *AccountSettings) GetEnabled1V1T() bool {
	if m != nil {
		return m.Enabled1V1T
	}
	return false
}

type DecisionResponse struct {
	//*
	// visitor ID
	//
	// Required
	//
	// The custom or decision generated visitor ID.
	VisitorId *wrappers.StringValue `protobuf:"bytes,1,opt,name=visitor_id,json=visitorId,proto3" json:"visitor_id,omitempty"`
	//*
	// decision group
	//
	// Optional
	//
	// The custom visitor group
	Campaigns []*Campaign `protobuf:"bytes,2,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
	//*
	// account settings
	//
	// Optional
	//
	// Account settings enabled and activated in the platform
	AccountSettings      *AccountSettings `protobuf:"bytes,3,opt,name=accountSettings,proto3" json:"accountSettings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DecisionResponse) Reset()         { *m = DecisionResponse{} }
func (m *DecisionResponse) String() string { return proto.CompactTextString(m) }
func (*DecisionResponse) ProtoMessage()    {}
func (*DecisionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{7}
}

func (m *DecisionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecisionResponse.Unmarshal(m, b)
}
func (m *DecisionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecisionResponse.Marshal(b, m, deterministic)
}
func (m *DecisionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecisionResponse.Merge(m, src)
}
func (m *DecisionResponse) XXX_Size() int {
	return xxx_messageInfo_DecisionResponse.Size(m)
}
func (m *DecisionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecisionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecisionResponse proto.InternalMessageInfo

func (m *DecisionResponse) GetVisitorId() *wrappers.StringValue {
	if m != nil {
		return m.VisitorId
	}
	return nil
}

func (m *DecisionResponse) GetCampaigns() []*Campaign {
	if m != nil {
		return m.Campaigns
	}
	return nil
}

func (m *DecisionResponse) GetAccountSettings() *AccountSettings {
	if m != nil {
		return m.AccountSettings
	}
	return nil
}

type DecisionResponsePanic struct {
	//*
	// visitor ID
	//
	// Required
	//
	// The custom or decision generated visitor ID.
	VisitorId *wrappers.StringValue `protobuf:"bytes,1,opt,name=visitor_id,json=visitorId,proto3" json:"visitor_id,omitempty"`
	//*
	// decision group
	//
	// Optional
	//
	// The custom visitor group
	Campaigns []*Campaign `protobuf:"bytes,2,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
	//*
	// panic
	//
	// Optional
	//
	// True if panic mode is set for the account environment
	Panic                bool     `protobuf:"varint,3,opt,name=panic,proto3" json:"panic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecisionResponsePanic) Reset()         { *m = DecisionResponsePanic{} }
func (m *DecisionResponsePanic) String() string { return proto.CompactTextString(m) }
func (*DecisionResponsePanic) ProtoMessage()    {}
func (*DecisionResponsePanic) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{8}
}

func (m *DecisionResponsePanic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecisionResponsePanic.Unmarshal(m, b)
}
func (m *DecisionResponsePanic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecisionResponsePanic.Marshal(b, m, deterministic)
}
func (m *DecisionResponsePanic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecisionResponsePanic.Merge(m, src)
}
func (m *DecisionResponsePanic) XXX_Size() int {
	return xxx_messageInfo_DecisionResponsePanic.Size(m)
}
func (m *DecisionResponsePanic) XXX_DiscardUnknown() {
	xxx_messageInfo_DecisionResponsePanic.DiscardUnknown(m)
}

var xxx_messageInfo_DecisionResponsePanic proto.InternalMessageInfo

func (m *DecisionResponsePanic) GetVisitorId() *wrappers.StringValue {
	if m != nil {
		return m.VisitorId
	}
	return nil
}

func (m *DecisionResponsePanic) GetCampaigns() []*Campaign {
	if m != nil {
		return m.Campaigns
	}
	return nil
}

func (m *DecisionResponsePanic) GetPanic() bool {
	if m != nil {
		return m.Panic
	}
	return false
}

type DecisionResponseFull struct {
	//*
	// visitor ID
	//
	// Required
	//
	// The custom or decision generated visitor ID.
	VisitorId *wrappers.StringValue `protobuf:"bytes,1,opt,name=visitor_id,json=visitorId,proto3" json:"visitor_id,omitempty"`
	//*
	// decision group
	//
	// Optional
	//
	// The custom visitor group
	Campaigns []*Campaign `protobuf:"bytes,2,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
	//*
	// campaigns_variation
	//
	// Optional
	//
	// An array of all the campaign and associated variation id
	CampaignsVariation []*CampaignIdVariationId `protobuf:"bytes,3,rep,name=campaigns_variation,json=campaignsVariation,proto3" json:"campaigns_variation,omitempty"`
	//*
	// merged_modifications
	//
	// Optional
	//
	// The merged modifications for all the campaigns
	MergedModifications *_struct.Struct `protobuf:"bytes,4,opt,name=merged_modifications,json=mergedModifications,proto3" json:"merged_modifications,omitempty"`
	//*
	// panic
	//
	// Optional
	//
	// True if panic mode is set for the account environment
	Panic                bool     `protobuf:"varint,5,opt,name=panic,proto3" json:"panic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecisionResponseFull) Reset()         { *m = DecisionResponseFull{} }
func (m *DecisionResponseFull) String() string { return proto.CompactTextString(m) }
func (*DecisionResponseFull) ProtoMessage()    {}
func (*DecisionResponseFull) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{9}
}

func (m *DecisionResponseFull) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecisionResponseFull.Unmarshal(m, b)
}
func (m *DecisionResponseFull) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecisionResponseFull.Marshal(b, m, deterministic)
}
func (m *DecisionResponseFull) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecisionResponseFull.Merge(m, src)
}
func (m *DecisionResponseFull) XXX_Size() int {
	return xxx_messageInfo_DecisionResponseFull.Size(m)
}
func (m *DecisionResponseFull) XXX_DiscardUnknown() {
	xxx_messageInfo_DecisionResponseFull.DiscardUnknown(m)
}

var xxx_messageInfo_DecisionResponseFull proto.InternalMessageInfo

func (m *DecisionResponseFull) GetVisitorId() *wrappers.StringValue {
	if m != nil {
		return m.VisitorId
	}
	return nil
}

func (m *DecisionResponseFull) GetCampaigns() []*Campaign {
	if m != nil {
		return m.Campaigns
	}
	return nil
}

func (m *DecisionResponseFull) GetCampaignsVariation() []*CampaignIdVariationId {
	if m != nil {
		return m.CampaignsVariation
	}
	return nil
}

func (m *DecisionResponseFull) GetMergedModifications() *_struct.Struct {
	if m != nil {
		return m.MergedModifications
	}
	return nil
}

func (m *DecisionResponseFull) GetPanic() bool {
	if m != nil {
		return m.Panic
	}
	return false
}

type DecisionResponseSimple struct {
	//*
	// campaigns_variation
	//
	// Optional
	//
	// An array of all the campaign and associated variation id
	CampaignsVariation []*CampaignIdVariationId `protobuf:"bytes,1,rep,name=campaigns_variation,json=campaignsVariation,proto3" json:"campaigns_variation,omitempty"`
	//*
	// merged_modifications
	//
	// Optional
	//
	// The merged modifications for all the campaigns
	MergedModifications  *_struct.Struct `protobuf:"bytes,2,opt,name=merged_modifications,json=mergedModifications,proto3" json:"merged_modifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DecisionResponseSimple) Reset()         { *m = DecisionResponseSimple{} }
func (m *DecisionResponseSimple) String() string { return proto.CompactTextString(m) }
func (*DecisionResponseSimple) ProtoMessage()    {}
func (*DecisionResponseSimple) Descriptor() ([]byte, []int) {
	return fileDescriptor_f38fa95d86323f29, []int{10}
}

func (m *DecisionResponseSimple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecisionResponseSimple.Unmarshal(m, b)
}
func (m *DecisionResponseSimple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecisionResponseSimple.Marshal(b, m, deterministic)
}
func (m *DecisionResponseSimple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecisionResponseSimple.Merge(m, src)
}
func (m *DecisionResponseSimple) XXX_Size() int {
	return xxx_messageInfo_DecisionResponseSimple.Size(m)
}
func (m *DecisionResponseSimple) XXX_DiscardUnknown() {
	xxx_messageInfo_DecisionResponseSimple.DiscardUnknown(m)
}

var xxx_messageInfo_DecisionResponseSimple proto.InternalMessageInfo

func (m *DecisionResponseSimple) GetCampaignsVariation() []*CampaignIdVariationId {
	if m != nil {
		return m.CampaignsVariation
	}
	return nil
}

func (m *DecisionResponseSimple) GetMergedModifications() *_struct.Struct {
	if m != nil {
		return m.MergedModifications
	}
	return nil
}

func init() {
	proto.RegisterEnum("flagship.protobuf.ModificationsType", ModificationsType_name, ModificationsType_value)
	proto.RegisterType((*CampaignIdVariationId)(nil), "flagship.protobuf.CampaignIdVariationId")
	proto.RegisterType((*Modifications)(nil), "flagship.protobuf.Modifications")
	proto.RegisterType((*TuringEngineOptions)(nil), "flagship.protobuf.TuringEngineOptions")
	proto.RegisterMapType((map[string]string)(nil), "flagship.protobuf.TuringEngineOptions.ContextKeysEntry")
	proto.RegisterType((*Variation)(nil), "flagship.protobuf.Variation")
	proto.RegisterType((*FullVariation)(nil), "flagship.protobuf.FullVariation")
	proto.RegisterType((*Campaign)(nil), "flagship.protobuf.Campaign")
	proto.RegisterType((*AccountSettings)(nil), "flagship.protobuf.AccountSettings")
	proto.RegisterType((*DecisionResponse)(nil), "flagship.protobuf.DecisionResponse")
	proto.RegisterType((*DecisionResponsePanic)(nil), "flagship.protobuf.DecisionResponsePanic")
	proto.RegisterType((*DecisionResponseFull)(nil), "flagship.protobuf.DecisionResponseFull")
	proto.RegisterType((*DecisionResponseSimple)(nil), "flagship.protobuf.DecisionResponseSimple")
}

func init() {
	proto.RegisterFile("decision_response/decision_response.proto", fileDescriptor_f38fa95d86323f29)
}

var fileDescriptor_f38fa95d86323f29 = []byte{
	// 806 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc6, 0x4e, 0x52, 0xe2, 0xe3, 0x0d, 0xf5, 0xce, 0xa6, 0x10, 0x2d, 0x51, 0x09, 0x16, 0x42,
	0x01, 0xd1, 0x44, 0x0d, 0x17, 0xfd, 0x41, 0x42, 0x5a, 0xb6, 0xde, 0xe0, 0x25, 0x9b, 0x56, 0x13,
	0x77, 0xb5, 0x94, 0x0b, 0xcb, 0xb1, 0x27, 0xee, 0xa8, 0x8e, 0xc7, 0xf8, 0x27, 0x6c, 0x5e, 0x81,
	0x97, 0xe0, 0x8a, 0x07, 0xe0, 0x09, 0xe0, 0x9e, 0x17, 0xe0, 0x71, 0x90, 0x7f, 0x62, 0x27, 0x4e,
	0xd8, 0x9f, 0x0b, 0xb4, 0x77, 0x33, 0x67, 0xbe, 0xef, 0xe4, 0x7c, 0xdf, 0xf8, 0x9c, 0x09, 0x7c,
	0x61, 0x11, 0x93, 0x06, 0x94, 0xb9, 0xba, 0x4f, 0x02, 0x8f, 0xb9, 0x01, 0xe9, 0x6f, 0x45, 0x7a,
	0x9e, 0xcf, 0x42, 0x86, 0xf6, 0x67, 0x8e, 0x61, 0x07, 0x6f, 0xa9, 0x97, 0xee, 0xa7, 0xd1, 0xec,
	0xf0, 0xa1, 0xcd, 0x98, 0xed, 0x90, 0xfe, 0x2a, 0xd0, 0xff, 0xc5, 0x37, 0x3c, 0x8f, 0xf8, 0x41,
	0x0a, 0x39, 0x6c, 0x97, 0xcf, 0x83, 0xd0, 0x8f, 0xcc, 0x30, 0x3d, 0x95, 0x7f, 0x82, 0x07, 0xc7,
	0xc6, 0xdc, 0x33, 0xa8, 0xed, 0xaa, 0xd6, 0xb9, 0xe1, 0x53, 0x23, 0xa4, 0xcc, 0x55, 0x2d, 0xf4,
	0x09, 0x88, 0x66, 0x76, 0xa0, 0x53, 0xab, 0xc5, 0x75, 0xb8, 0xae, 0x80, 0xc1, 0xcc, 0xb1, 0xe8,
	0x53, 0xd8, 0x5b, 0xac, 0xf0, 0x31, 0x82, 0x4f, 0x10, 0xe2, 0xa2, 0xc8, 0x21, 0x5f, 0x42, 0xe3,
	0x8c, 0x59, 0x74, 0x46, 0xcd, 0x24, 0x12, 0xa0, 0xa7, 0x50, 0x0d, 0x97, 0x1e, 0x49, 0xb2, 0x7d,
	0x30, 0xf8, 0xac, 0xb7, 0xa5, 0xa6, 0xb7, 0x81, 0xd7, 0x96, 0x1e, 0xc1, 0x09, 0x03, 0x3d, 0x82,
	0xda, 0xc2, 0x70, 0x22, 0x92, 0xfc, 0x8c, 0x38, 0xf8, 0xa8, 0x97, 0xaa, 0x2a, 0x88, 0x93, 0x44,
	0x15, 0x4e, 0x51, 0xf2, 0x1f, 0x1c, 0x1c, 0x68, 0x91, 0x4f, 0x5d, 0x5b, 0x71, 0x6d, 0xea, 0x92,
	0x97, 0x5e, 0x5a, 0xc0, 0x1b, 0xd8, 0x33, 0x99, 0x1b, 0x92, 0xcb, 0x50, 0x7f, 0x47, 0x96, 0x41,
	0x8b, 0xeb, 0x54, 0xba, 0xe2, 0xe0, 0xc9, 0x8e, 0x42, 0x76, 0xb0, 0x7b, 0xc7, 0x29, 0xf5, 0x07,
	0xb2, 0x0c, 0x14, 0x37, 0xf4, 0x97, 0x58, 0x34, 0x8b, 0xc8, 0xe1, 0xb7, 0x20, 0x95, 0x01, 0x48,
	0x82, 0xca, 0x3b, 0xb2, 0xcc, 0xdc, 0x8b, 0x97, 0xa8, 0xb9, 0x2e, 0x44, 0xc8, 0xea, 0x7d, 0xce,
	0x3f, 0xe5, 0xe4, 0xdf, 0x38, 0x10, 0xf2, 0x1b, 0x40, 0x5f, 0x01, 0x9f, 0xd9, 0x2e, 0x0e, 0xda,
	0xbb, 0xd4, 0x52, 0xd7, 0x3e, 0x8f, 0xb9, 0x98, 0xa7, 0x16, 0x3a, 0x81, 0xc6, 0x7c, 0xdd, 0xb9,
	0xcc, 0xa6, 0xce, 0x75, 0x0e, 0xe3, 0x4d, 0x1a, 0x6a, 0x83, 0xe0, 0x93, 0x19, 0xf1, 0x89, 0x6b,
	0x92, 0x56, 0xa5, 0xc3, 0x75, 0xeb, 0xb8, 0x08, 0xc8, 0xbf, 0xf2, 0xd0, 0x38, 0x89, 0x1c, 0xe7,
	0xae, 0xab, 0x54, 0x00, 0x42, 0xa2, 0xb3, 0xf4, 0x56, 0x92, 0x32, 0xc5, 0xc1, 0xe7, 0x37, 0xbb,
	0x43, 0x2c, 0x84, 0xf9, 0xc7, 0xf0, 0x10, 0xc0, 0x70, 0x1c, 0x96, 0x66, 0x6d, 0x55, 0x3b, 0x5c,
	0xb7, 0x86, 0xd7, 0x22, 0x9b, 0x66, 0xd4, 0xca, 0x66, 0xfc, 0xc5, 0x41, 0x7d, 0xd5, 0x3a, 0xb7,
	0xf4, 0xe1, 0x14, 0x50, 0xd1, 0x3a, 0xb6, 0xcf, 0x22, 0x6f, 0xd5, 0x40, 0xd7, 0xb1, 0xa5, 0x9c,
	0x37, 0x8c, 0x69, 0xaa, 0x85, 0x9e, 0x83, 0x90, 0xc7, 0x32, 0x2b, 0xda, 0x3b, 0xac, 0xc8, 0xaf,
	0x0c, 0x17, 0x70, 0x79, 0x02, 0xf7, 0x8f, 0x4c, 0x93, 0x45, 0x6e, 0x38, 0x21, 0x61, 0x48, 0x5d,
	0x3b, 0xf1, 0x84, 0xb8, 0xc6, 0xd4, 0x21, 0xd6, 0xc5, 0xab, 0x8b, 0x44, 0x50, 0x1d, 0xaf, 0x45,
	0x50, 0x07, 0xc4, 0x6c, 0xf7, 0xf8, 0xfc, 0xb1, 0x96, 0xd4, 0x5c, 0xc7, 0xeb, 0x21, 0xf9, 0x1f,
	0x0e, 0xa4, 0x17, 0xd9, 0xf8, 0xc2, 0xd9, 0xf4, 0x42, 0xdf, 0x00, 0x2c, 0x68, 0x40, 0x43, 0xe6,
	0xeb, 0x37, 0xf4, 0x49, 0xc8, 0xf0, 0xaa, 0x85, 0x9e, 0x81, 0xb0, 0x9a, 0x3b, 0xf1, 0x27, 0x13,
	0x77, 0xec, 0xc7, 0x3b, 0x24, 0xae, 0x2e, 0x03, 0x17, 0x68, 0x34, 0x82, 0xfb, 0xc6, 0xa6, 0xc2,
	0xcc, 0x23, 0x79, 0x47, 0x82, 0x92, 0x17, 0xb8, 0x4c, 0x95, 0x7f, 0xe7, 0xe0, 0x41, 0x59, 0xda,
	0x2b, 0xc3, 0xa5, 0xe6, 0x9d, 0xe9, 0x6b, 0x42, 0xcd, 0x8b, 0x0b, 0xc8, 0x7a, 0x35, 0xdd, 0xc8,
	0x7f, 0xf3, 0xd0, 0x2c, 0xd7, 0x19, 0xf7, 0xed, 0x9d, 0x95, 0xf9, 0x23, 0x1c, 0xe4, 0x1b, 0x7d,
	0xfd, 0x73, 0x8d, 0x93, 0x74, 0xaf, 0x48, 0xb2, 0xf1, 0x26, 0x61, 0x94, 0x27, 0x29, 0x26, 0xd0,
	0x29, 0x34, 0xe7, 0xc4, 0xb7, 0x89, 0xa5, 0x6f, 0x8e, 0x96, 0xea, 0xd5, 0xef, 0xc4, 0x41, 0x4a,
	0xda, 0x7c, 0x9e, 0x72, 0x37, 0x6b, 0xeb, 0x6e, 0xfe, 0xc9, 0xc1, 0x87, 0x65, 0x37, 0x27, 0x74,
	0xee, 0x39, 0xe4, 0xbf, 0x74, 0x71, 0xff, 0xa3, 0x2e, 0xfe, 0xf6, 0xba, 0xbe, 0xfc, 0x19, 0xf6,
	0xb7, 0xde, 0x55, 0x54, 0x87, 0xea, 0xf8, 0xf5, 0x68, 0x24, 0xbd, 0x17, 0xaf, 0x4e, 0x27, 0x2f,
	0xc7, 0x12, 0x17, 0xaf, 0x34, 0xe5, 0x42, 0x93, 0x78, 0x24, 0x40, 0x4d, 0x3d, 0x3b, 0x1a, 0x2a,
	0x52, 0x25, 0x0e, 0x7e, 0xaf, 0x9d, 0x8d, 0xa4, 0x6a, 0xbc, 0x3a, 0x19, 0x1d, 0x0d, 0xa5, 0x1a,
	0xda, 0x83, 0x3a, 0x56, 0x5e, 0xa8, 0x58, 0x39, 0xd6, 0xa4, 0x7b, 0x68, 0x1f, 0x1a, 0xda, 0x6b,
	0xac, 0x8e, 0x87, 0xba, 0x32, 0x1e, 0xaa, 0x63, 0x45, 0x7a, 0xff, 0xbb, 0x67, 0x6f, 0x9e, 0xd8,
	0x34, 0x7c, 0x1b, 0x4d, 0x7b, 0x26, 0x9b, 0xf7, 0x57, 0x46, 0x3c, 0xa2, 0xac, 0x58, 0x27, 0xa5,
	0x6f, 0xff, 0xd3, 0x99, 0xde, 0x4b, 0x0e, 0xbe, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xb4,
	0x61, 0xb7, 0x17, 0x09, 0x00, 0x00,
}
