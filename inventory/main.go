package main

import (
	"flag"
	"fmt"
	"ha-network-service/grpcpb/inventory"
	"ha-network-service/inventory/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8100, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	inventory.RegisterInventoryServer(grpcServer, server.NewInventoryServer())
	grpcServer.Serve(lis)
}
