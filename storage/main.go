package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/eddielin0926/ha-network-service/grpcpb/storage"
	"github.com/eddielin0926/ha-network-service/storage/models"
	"github.com/eddielin0926/ha-network-service/storage/server"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbuser := os.Getenv("DB_USERNAME")
	dbpasswd := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_ADDRESS")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpasswd, dbhost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect database")
	}
	db.AutoMigrate(&models.Record{})

	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	storage.RegisterStorageServer(grpcServer, server.NewStorageServer(db))
	grpcServer.Serve(lis)
}
