package main

import (
	"flag"
	"fmt"
	"ha-network-service/grpcpb/storage"
	"ha-network-service/storage/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8200, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	storage.RegisterStorageServer(grpcServer, server.NewStorageServer())
	grpcServer.Serve(lis)
}