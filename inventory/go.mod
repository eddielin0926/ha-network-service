module ha-network-service/inventory

go 1.19

require (
	google.golang.org/grpc v1.52.3
	ha-network-service/grpcpb v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace ha-network-service/grpcpb => ../grpcpb