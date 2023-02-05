package main

import (
	"flag"
	"ha-network-service/business/initialize"
	"ha-network-service/business/routes"
	"ha-network-service/grpcpb/inventory"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8100", "the address to connect to")
)

func init() {
	initialize.LoadEnv()
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := inventory.NewInventoryClient(conn)

	app := gin.Default()

	route := routes.NewApiRoute(c)
	route.Setup(app)

	app.Run()
}
