package main

import (
	"flag"
	"log"

	"github.com/eddielin0926/ha-network-service/business/initialize"
	"github.com/eddielin0926/ha-network-service/business/routes"
	"github.com/eddielin0926/ha-network-service/grpcpb/inventory"
	"github.com/eddielin0926/ha-network-service/grpcpb/storage"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	inventoryUrl = flag.String("inv", "localhost:8100", "the inventory address for connection")
	storageUrl   = flag.String("stor", "localhost:8200", "the storage address for connection")
)

func init() {
	initialize.LoadEnv()
}

func main() {
	// CLI flag
	flag.Parse()

	// Inventory gRPC
	invConn, err := grpc.Dial(*inventoryUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer invConn.Close()
	invc := inventory.NewInventoryClient(invConn)

	// Storage gRPC
	storConn, err := grpc.Dial(*storageUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer storConn.Close()
	storc := storage.NewStorageClient(storConn)

	// Gin
	app := gin.Default()
	route := routes.NewApiRoute(invc, storc)
	route.Setup(app)
	app.Run()
}
