// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: currency.proto

package currency

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Currencies int32

const (
	Currencies_EUR Currencies = 0
	Currencies_USD Currencies = 1
	Currencies_JPY Currencies = 2
	Currencies_BGN Currencies = 3
	Currencies_CZK Currencies = 4
	Currencies_DKK Currencies = 5
	Currencies_GBP Currencies = 6
	Currencies_HUF Currencies = 7
	Currencies_PLN Currencies = 8
	Currencies_RON Currencies = 9
	Currencies_SEK Currencies = 10
	Currencies_CHF Currencies = 11
	Currencies_ISK Currencies = 12
	Currencies_NOK Currencies = 13
	Currencies_HRK Currencies = 14
	Currencies_RUB Currencies = 15
	Currencies_TRY Currencies = 16
	Currencies_AUD Currencies = 17
	Currencies_BRL Currencies = 18
	Currencies_CAD Currencies = 19
	Currencies_CNY Currencies = 20
	Currencies_HKD Currencies = 21
	Currencies_IDR Currencies = 22
	Currencies_ILS Currencies = 23
	Currencies_INR Currencies = 24
	Currencies_KRW Currencies = 25
	Currencies_MXN Currencies = 26
	Currencies_MYR Currencies = 27
	Currencies_NZD Currencies = 28
	Currencies_PHP Currencies = 29
	Currencies_SGD Currencies = 30
	Currencies_THB Currencies = 31
	Currencies_ZAR Currencies = 32
)

// Enum value maps for Currencies.
var (
	Currencies_name = map[int32]string{
		0:  "EUR",
		1:  "USD",
		2:  "JPY",
		3:  "BGN",
		4:  "CZK",
		5:  "DKK",
		6:  "GBP",
		7:  "HUF",
		8:  "PLN",
		9:  "RON",
		10: "SEK",
		11: "CHF",
		12: "ISK",
		13: "NOK",
		14: "HRK",
		15: "RUB",
		16: "TRY",
		17: "AUD",
		18: "BRL",
		19: "CAD",
		20: "CNY",
		21: "HKD",
		22: "IDR",
		23: "ILS",
		24: "INR",
		25: "KRW",
		26: "MXN",
		27: "MYR",
		28: "NZD",
		29: "PHP",
		30: "SGD",
		31: "THB",
		32: "ZAR",
	}
	Currencies_value = map[string]int32{
		"EUR": 0,
		"USD": 1,
		"JPY": 2,
		"BGN": 3,
		"CZK": 4,
		"DKK": 5,
		"GBP": 6,
		"HUF": 7,
		"PLN": 8,
		"RON": 9,
		"SEK": 10,
		"CHF": 11,
		"ISK": 12,
		"NOK": 13,
		"HRK": 14,
		"RUB": 15,
		"TRY": 16,
		"AUD": 17,
		"BRL": 18,
		"CAD": 19,
		"CNY": 20,
		"HKD": 21,
		"IDR": 22,
		"ILS": 23,
		"INR": 24,
		"KRW": 25,
		"MXN": 26,
		"MYR": 27,
		"NZD": 28,
		"PHP": 29,
		"SGD": 30,
		"THB": 31,
		"ZAR": 32,
	}
)

func (x Currencies) Enum() *Currencies {
	p := new(Currencies)
	*p = x
	return p
}

func (x Currencies) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currencies) Descriptor() protoreflect.EnumDescriptor {
	return file_currency_proto_enumTypes[0].Descriptor()
}

func (Currencies) Type() protoreflect.EnumType {
	return &file_currency_proto_enumTypes[0]
}

func (x Currencies) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currencies.Descriptor instead.
func (Currencies) EnumDescriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{0}
}

type RateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base        Currencies `protobuf:"varint,1,opt,name=Base,proto3,enum=Currencies" json:"Base,omitempty"`
	Destination Currencies `protobuf:"varint,2,opt,name=Destination,proto3,enum=Currencies" json:"Destination,omitempty"`
}

func (x *RateRequest) Reset() {
	*x = RateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRequest) ProtoMessage() {}

func (x *RateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRequest.ProtoReflect.Descriptor instead.
func (*RateRequest) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{0}
}

func (x *RateRequest) GetBase() Currencies {
	if x != nil {
		return x.Base
	}
	return Currencies_EUR
}

func (x *RateRequest) GetDestination() Currencies {
	if x != nil {
		return x.Destination
	}
	return Currencies_EUR
}

type RateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base        Currencies `protobuf:"varint,1,opt,name=Base,proto3,enum=Currencies" json:"Base,omitempty"`
	Destination Currencies `protobuf:"varint,2,opt,name=Destination,proto3,enum=Currencies" json:"Destination,omitempty"`
	Rate        float64    `protobuf:"fixed64,3,opt,name=Rate,proto3" json:"Rate,omitempty"`
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{1}
}

func (x *RateResponse) GetBase() Currencies {
	if x != nil {
		return x.Base
	}
	return Currencies_EUR
}

func (x *RateResponse) GetDestination() Currencies {
	if x != nil {
		return x.Destination
	}
	return Currencies_EUR
}

func (x *RateResponse) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type StreamingRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*StreamingRateResponse_RateResponse
	//	*StreamingRateResponse_Error
	Message isStreamingRateResponse_Message `protobuf_oneof:"message"`
}

func (x *StreamingRateResponse) Reset() {
	*x = StreamingRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingRateResponse) ProtoMessage() {}

func (x *StreamingRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingRateResponse.ProtoReflect.Descriptor instead.
func (*StreamingRateResponse) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{2}
}

func (m *StreamingRateResponse) GetMessage() isStreamingRateResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *StreamingRateResponse) GetRateResponse() *RateResponse {
	if x, ok := x.GetMessage().(*StreamingRateResponse_RateResponse); ok {
		return x.RateResponse
	}
	return nil
}

func (x *StreamingRateResponse) GetError() *status.Status {
	if x, ok := x.GetMessage().(*StreamingRateResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isStreamingRateResponse_Message interface {
	isStreamingRateResponse_Message()
}

type StreamingRateResponse_RateResponse struct {
	RateResponse *RateResponse `protobuf:"bytes,1,opt,name=rate_response,json=rateResponse,proto3,oneof"`
}

type StreamingRateResponse_Error struct {
	Error *status.Status `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*StreamingRateResponse_RateResponse) isStreamingRateResponse_Message() {}

func (*StreamingRateResponse_Error) isStreamingRateResponse_Message() {}

var File_currency_proto protoreflect.FileDescriptor

var file_currency_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0b, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x72, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x52, 0x61, 0x74, 0x65, 0x22, 0x84, 0x01, 0x0a,
	0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0xb5, 0x02, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x55, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55,
	0x53, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x50, 0x59, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x42, 0x47, 0x4e, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x5a, 0x4b, 0x10, 0x04, 0x12,
	0x07, 0x0a, 0x03, 0x44, 0x4b, 0x4b, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x42, 0x50, 0x10,
	0x06, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x55, 0x46, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4c,
	0x4e, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x07, 0x0a, 0x03,
	0x53, 0x45, 0x4b, 0x10, 0x0a, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x48, 0x46, 0x10, 0x0b, 0x12, 0x07,
	0x0a, 0x03, 0x49, 0x53, 0x4b, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x4f, 0x4b, 0x10, 0x0d,
	0x12, 0x07, 0x0a, 0x03, 0x48, 0x52, 0x4b, 0x10, 0x0e, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x55, 0x42,
	0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x52, 0x59, 0x10, 0x10, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x55, 0x44, 0x10, 0x11, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x52, 0x4c, 0x10, 0x12, 0x12, 0x07, 0x0a,
	0x03, 0x43, 0x41, 0x44, 0x10, 0x13, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x4e, 0x59, 0x10, 0x14, 0x12,
	0x07, 0x0a, 0x03, 0x48, 0x4b, 0x44, 0x10, 0x15, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x44, 0x52, 0x10,
	0x16, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4c, 0x53, 0x10, 0x17, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4e,
	0x52, 0x10, 0x18, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x52, 0x57, 0x10, 0x19, 0x12, 0x07, 0x0a, 0x03,
	0x4d, 0x58, 0x4e, 0x10, 0x1a, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x59, 0x52, 0x10, 0x1b, 0x12, 0x07,
	0x0a, 0x03, 0x4e, 0x5a, 0x44, 0x10, 0x1c, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x48, 0x50, 0x10, 0x1d,
	0x12, 0x07, 0x0a, 0x03, 0x53, 0x47, 0x44, 0x10, 0x1e, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x48, 0x42,
	0x10, 0x1f, 0x12, 0x07, 0x0a, 0x03, 0x5a, 0x41, 0x52, 0x10, 0x20, 0x32, 0x6e, 0x0a, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x26, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x0c, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x0c, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_currency_proto_rawDescOnce sync.Once
	file_currency_proto_rawDescData = file_currency_proto_rawDesc
)

func file_currency_proto_rawDescGZIP() []byte {
	file_currency_proto_rawDescOnce.Do(func() {
		file_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_currency_proto_rawDescData)
	})
	return file_currency_proto_rawDescData
}

var file_currency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_currency_proto_goTypes = []interface{}{
	(Currencies)(0),               // 0: Currencies
	(*RateRequest)(nil),           // 1: RateRequest
	(*RateResponse)(nil),          // 2: RateResponse
	(*StreamingRateResponse)(nil), // 3: StreamingRateResponse
	(*status.Status)(nil),         // 4: google.rpc.Status
}
var file_currency_proto_depIdxs = []int32{
	0, // 0: RateRequest.Base:type_name -> Currencies
	0, // 1: RateRequest.Destination:type_name -> Currencies
	0, // 2: RateResponse.Base:type_name -> Currencies
	0, // 3: RateResponse.Destination:type_name -> Currencies
	2, // 4: StreamingRateResponse.rate_response:type_name -> RateResponse
	4, // 5: StreamingRateResponse.error:type_name -> google.rpc.Status
	1, // 6: Currency.GetRate:input_type -> RateRequest
	1, // 7: Currency.SubscribeRates:input_type -> RateRequest
	2, // 8: Currency.GetRate:output_type -> RateResponse
	3, // 9: Currency.SubscribeRates:output_type -> StreamingRateResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_currency_proto_init() }
func file_currency_proto_init() {
	if File_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateRequest); i {
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
		file_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResponse); i {
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
		file_currency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingRateResponse); i {
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
	file_currency_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*StreamingRateResponse_RateResponse)(nil),
		(*StreamingRateResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_currency_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currency_proto_goTypes,
		DependencyIndexes: file_currency_proto_depIdxs,
		EnumInfos:         file_currency_proto_enumTypes,
		MessageInfos:      file_currency_proto_msgTypes,
	}.Build()
	File_currency_proto = out.File
	file_currency_proto_rawDesc = nil
	file_currency_proto_goTypes = nil
	file_currency_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyClient interface {
	GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error)
	SubscribeRates(ctx context.Context, opts ...grpc.CallOption) (Currency_SubscribeRatesClient, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetRate(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyClient) SubscribeRates(ctx context.Context, opts ...grpc.CallOption) (Currency_SubscribeRatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Currency_serviceDesc.Streams[0], "/Currency/SubscribeRates", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencySubscribeRatesClient{stream}
	return x, nil
}

type Currency_SubscribeRatesClient interface {
	Send(*RateRequest) error
	Recv() (*StreamingRateResponse, error)
	grpc.ClientStream
}

type currencySubscribeRatesClient struct {
	grpc.ClientStream
}

func (x *currencySubscribeRatesClient) Send(m *RateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *currencySubscribeRatesClient) Recv() (*StreamingRateResponse, error) {
	m := new(StreamingRateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyServer is the server API for Currency service.
type CurrencyServer interface {
	GetRate(context.Context, *RateRequest) (*RateResponse, error)
	SubscribeRates(Currency_SubscribeRatesServer) error
}

// UnimplementedCurrencyServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (*UnimplementedCurrencyServer) GetRate(context.Context, *RateRequest) (*RateResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetRate not implemented")
}
func (*UnimplementedCurrencyServer) SubscribeRates(Currency_SubscribeRatesServer) error {
	return status1.Errorf(codes.Unimplemented, "method SubscribeRates not implemented")
}

func RegisterCurrencyServer(s *grpc.Server, srv CurrencyServer) {
	s.RegisterService(&_Currency_serviceDesc, srv)
}

func _Currency_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetRate(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Currency_SubscribeRates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CurrencyServer).SubscribeRates(&currencySubscribeRatesServer{stream})
}

type Currency_SubscribeRatesServer interface {
	Send(*StreamingRateResponse) error
	Recv() (*RateRequest, error)
	grpc.ServerStream
}

type currencySubscribeRatesServer struct {
	grpc.ServerStream
}

func (x *currencySubscribeRatesServer) Send(m *StreamingRateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *currencySubscribeRatesServer) Recv() (*RateRequest, error) {
	m := new(RateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Currency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _Currency_GetRate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeRates",
			Handler:       _Currency_SubscribeRates_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "currency.proto",
}
