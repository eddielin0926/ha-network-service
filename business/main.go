package main

import (
	"fmt"
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
	invaddr := os.Getenv("INVENTORY_ADDRESS")
	invport := os.Getenv("INVENTORY_PORT")
	storaddr := os.Getenv("STORAGE_ADDRESS")
	storport := os.Getenv("STORAGE_PORT")
	invurl := fmt.Sprintf("%s:%s", invaddr, invport)
	storurl := fmt.Sprintf("%s:%s", storaddr, storport)

	// Inventory gRPC
	invConn, err := grpc.Dial(invurl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer invConn.Close()
	invc := inventory.NewInventoryClient(invConn)

	// Storage gRPC
	storConn, err := grpc.Dial(storurl, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
