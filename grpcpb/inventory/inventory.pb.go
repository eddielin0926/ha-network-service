// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: grpcpb/inventory/inventory.proto

package inventory

import (
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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	C int32 `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D int32 `protobuf:"varint,4,opt,name=d,proto3" json:"d,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcpb_inventory_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_grpcpb_inventory_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_grpcpb_inventory_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Order) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *Order) GetC() int32 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *Order) GetD() int32 {
	if x != nil {
		return x.D
	}
	return 0
}

type Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Material  int32  `protobuf:"varint,1,opt,name=material,proto3" json:"material,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Material) Reset() {
	*x = Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcpb_inventory_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Material) ProtoMessage() {}

func (x *Material) ProtoReflect() protoreflect.Message {
	mi := &file_grpcpb_inventory_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Material.ProtoReflect.Descriptor instead.
func (*Material) Descriptor() ([]byte, []int) {
	return file_grpcpb_inventory_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *Material) GetMaterial() int32 {
	if x != nil {
		return x.Material
	}
	return 0
}

func (x *Material) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_grpcpb_inventory_inventory_proto protoreflect.FileDescriptor

var file_grpcpb_inventory_inventory_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x3f, 0x0a,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x63,
	0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x64, 0x22, 0x44,
	0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x32, 0x41, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x12, 0x10, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x1a, 0x13, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x42, 0x25, 0x5a, 0x23, 0x68, 0x61, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcpb_inventory_inventory_proto_rawDescOnce sync.Once
	file_grpcpb_inventory_inventory_proto_rawDescData = file_grpcpb_inventory_inventory_proto_rawDesc
)

func file_grpcpb_inventory_inventory_proto_rawDescGZIP() []byte {
	file_grpcpb_inventory_inventory_proto_rawDescOnce.Do(func() {
		file_grpcpb_inventory_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcpb_inventory_inventory_proto_rawDescData)
	})
	return file_grpcpb_inventory_inventory_proto_rawDescData
}

var file_grpcpb_inventory_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpcpb_inventory_inventory_proto_goTypes = []interface{}{
	(*Order)(nil),    // 0: inventory.Order
	(*Material)(nil), // 1: inventory.Material
}
var file_grpcpb_inventory_inventory_proto_depIdxs = []int32{
	0, // 0: inventory.Inventory.GetMaterial:input_type -> inventory.Order
	1, // 1: inventory.Inventory.GetMaterial:output_type -> inventory.Material
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcpb_inventory_inventory_proto_init() }
func file_grpcpb_inventory_inventory_proto_init() {
	if File_grpcpb_inventory_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcpb_inventory_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_grpcpb_inventory_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Material); i {
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
			RawDescriptor: file_grpcpb_inventory_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcpb_inventory_inventory_proto_goTypes,
		DependencyIndexes: file_grpcpb_inventory_inventory_proto_depIdxs,
		MessageInfos:      file_grpcpb_inventory_inventory_proto_msgTypes,
	}.Build()
	File_grpcpb_inventory_inventory_proto = out.File
	file_grpcpb_inventory_inventory_proto_rawDesc = nil
	file_grpcpb_inventory_inventory_proto_goTypes = nil
	file_grpcpb_inventory_inventory_proto_depIdxs = nil
}
