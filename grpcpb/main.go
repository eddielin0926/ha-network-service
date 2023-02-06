package grpcpb

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProtobufToJson(msg protoreflect.ProtoMessage) ([]byte, error) {
	m := protojson.MarshalOptions{
		UseEnumNumbers:  true,
		EmitUnpopulated: true,
	}
	return m.Marshal(msg)
}
