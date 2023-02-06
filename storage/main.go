package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/eddielin0926/ha-network-service/grpcpb/storage"
	"github.com/eddielin0926/ha-network-service/storage/server"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	storage.RegisterStorageServer(grpcServer, server.NewStorageServer())
	grpcServer.Serve(lis)
}
