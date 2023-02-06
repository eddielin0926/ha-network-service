package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/eddielin0926/ha-network-service/grpcpb/inventory"
	"github.com/eddielin0926/ha-network-service/inventory/server"

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
	inventory.RegisterInventoryServer(grpcServer, server.NewInventoryServer())
	grpcServer.Serve(lis)
}
