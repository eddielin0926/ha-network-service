// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: grpcpb/storage/storage.proto

package storage

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Status int32

const (
	Status_FAIL    Status = 0
	Status_SUCCESS Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
	}
	Status_value = map[string]int32{
		"FAIL":    0,
		"SUCCESS": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_grpcpb_storage_storage_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_grpcpb_storage_storage_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_grpcpb_storage_storage_proto_rawDescGZIP(), []int{0}
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location  string               `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature string               `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Material  float32              `protobuf:"fixed32,4,opt,name=material,proto3" json:"material,omitempty"`
	A         float32              `protobuf:"fixed32,5,opt,name=a,proto3" json:"a,omitempty"`
	B         float32              `protobuf:"fixed32,6,opt,name=b,proto3" json:"b,omitempty"`
	C         float32              `protobuf:"fixed32,7,opt,name=c,proto3" json:"c,omitempty"`
	D         float32              `protobuf:"fixed32,8,opt,name=d,proto3" json:"d,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcpb_storage_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_grpcpb_storage_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_grpcpb_storage_storage_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Record) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Record) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Record) GetMaterial() float32 {
	if x != nil {
		return x.Material
	}
	return 0
}

func (x *Record) GetA() float32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Record) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *Record) GetC() float32 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *Record) GetD() float32 {
	if x != nil {
		return x.D
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=storage.Status" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcpb_storage_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpcpb_storage_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpcpb_storage_storage_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_FAIL
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string               `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Date     *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcpb_storage_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_grpcpb_storage_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_grpcpb_storage_storage_proto_rawDescGZIP(), []int{2}
}

func (x *Query) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Query) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type RecordsRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *RecordsRespone) Reset() {
	*x = RecordsRespone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcpb_storage_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordsRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordsRespone) ProtoMessage() {}

func (x *RecordsRespone) ProtoReflect() protoreflect.Message {
	mi := &file_grpcpb_storage_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordsRespone.ProtoReflect.Descriptor instead.
func (*RecordsRespone) Descriptor() ([]byte, []int) {
	return file_grpcpb_storage_storage_proto_rawDescGZIP(), []int{3}
}

func (x *RecordsRespone) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string               `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Date     *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Count    int32                `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Material float32              `protobuf:"fixed32,4,opt,name=material,proto3" json:"material,omitempty"`
	A        float32              `protobuf:"fixed32,5,opt,name=a,proto3" json:"a,omitempty"`
	B        float32              `protobuf:"fixed32,6,opt,name=b,proto3" json:"b,omitempty"`
	C        float32              `protobuf:"fixed32,7,opt,name=c,proto3" json:"c,omitempty"`
	D        float32              `protobuf:"fixed32,8,opt,name=d,proto3" json:"d,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcpb_storage_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_grpcpb_storage_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_grpcpb_storage_storage_proto_rawDescGZIP(), []int{4}
}

func (x *Report) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Report) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Report) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Report) GetMaterial() float32 {
	if x != nil {
		return x.Material
	}
	return 0
}

func (x *Report) GetA() float32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Report) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *Report) GetC() float32 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *Report) GetD() float32 {
	if x != nil {
		return x.D
	}
	return 0
}

var File_grpcpb_storage_storage_proto protoreflect.FileDescriptor

var file_grpcpb_storage_storage_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x62, 0x12,
	0x0c, 0x0a, 0x01, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x63, 0x12, 0x0c, 0x0a,
	0x01, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x64, 0x22, 0x33, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x53, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x63, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x64, 0x2a, 0x1f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a,
	0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x32, 0xa0, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x30, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0f,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a,
	0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x23, 0x5a, 0x21, 0x68, 0x61, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcpb_storage_storage_proto_rawDescOnce sync.Once
	file_grpcpb_storage_storage_proto_rawDescData = file_grpcpb_storage_storage_proto_rawDesc
)

func file_grpcpb_storage_storage_proto_rawDescGZIP() []byte {
	file_grpcpb_storage_storage_proto_rawDescOnce.Do(func() {
		file_grpcpb_storage_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcpb_storage_storage_proto_rawDescData)
	})
	return file_grpcpb_storage_storage_proto_rawDescData
}

var file_grpcpb_storage_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpcpb_storage_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpcpb_storage_storage_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: storage.Status
	(*Record)(nil),              // 1: storage.Record
	(*Response)(nil),            // 2: storage.Response
	(*Query)(nil),               // 3: storage.Query
	(*RecordsRespone)(nil),      // 4: storage.RecordsRespone
	(*Report)(nil),              // 5: storage.Report
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_grpcpb_storage_storage_proto_depIdxs = []int32{
	6, // 0: storage.Record.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: storage.Response.status:type_name -> storage.Status
	6, // 2: storage.Query.date:type_name -> google.protobuf.Timestamp
	1, // 3: storage.RecordsRespone.records:type_name -> storage.Record
	6, // 4: storage.Report.date:type_name -> google.protobuf.Timestamp
	1, // 5: storage.Storage.SaveRecord:input_type -> storage.Record
	3, // 6: storage.Storage.GetRecords:input_type -> storage.Query
	3, // 7: storage.Storage.GetReport:input_type -> storage.Query
	2, // 8: storage.Storage.SaveRecord:output_type -> storage.Response
	4, // 9: storage.Storage.GetRecords:output_type -> storage.RecordsRespone
	5, // 10: storage.Storage.GetReport:output_type -> storage.Report
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpcpb_storage_storage_proto_init() }
func file_grpcpb_storage_storage_proto_init() {
	if File_grpcpb_storage_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcpb_storage_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_grpcpb_storage_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_grpcpb_storage_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_grpcpb_storage_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordsRespone); i {
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
		file_grpcpb_storage_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
			RawDescriptor: file_grpcpb_storage_storage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcpb_storage_storage_proto_goTypes,
		DependencyIndexes: file_grpcpb_storage_storage_proto_depIdxs,
		EnumInfos:         file_grpcpb_storage_storage_proto_enumTypes,
		MessageInfos:      file_grpcpb_storage_storage_proto_msgTypes,
	}.Build()
	File_grpcpb_storage_storage_proto = out.File
	file_grpcpb_storage_storage_proto_rawDesc = nil
	file_grpcpb_storage_storage_proto_goTypes = nil
	file_grpcpb_storage_storage_proto_depIdxs = nil
}
