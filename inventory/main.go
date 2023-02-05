package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	pb "ha-network-service/grpcpb/inventory"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8100, "The server port")
)

type inventoryServer struct {
	pb.UnimplementedInventoryServer
}

func (s *inventoryServer) GetMaterial(ctx context.Context, in *pb.Order) (*pb.Material, error) {
	a, b, c, d := in.GetA(), in.GetB(), in.GetC(), in.GetD()
	material := a*3 + b*2 + c*4 + d*10
	signature := s.sign(a, b, c, d)
	return &pb.Material{Material: material, Signature: signature}, nil
}

func (s *inventoryServer) sign(a, b, c, d int32) string {
	data := a + b + c + d
	strData := strconv.FormatInt(int64(data), 10)
	signature := base64.StdEncoding.EncodeToString([]byte(strData))
	return signature
}

func newServer() *inventoryServer {
	s := &inventoryServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterInventoryServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
