package main

import (
	"log"
	"os"

	"github.com/eddielin0926/ha-network-service/business/initialize"
	"github.com/eddielin0926/ha-network-service/business/routes"
	"github.com/eddielin0926/ha-network-service/grpcpb/inventory"
	"github.com/eddielin0926/ha-network-service/grpcpb/storage"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	initialize.LoadEnv()
}

func main() {
	inventoryUrl := os.Getenv("INVENTORY_URL")
	storageUrl := os.Getenv("STORAGE_URL")

	// Inventory gRPC
	invConn, err := grpc.Dial(inventoryUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer invConn.Close()
	invc := inventory.NewInventoryClient(invConn)

	// Storage gRPC
	storConn, err := grpc.Dial(storageUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
