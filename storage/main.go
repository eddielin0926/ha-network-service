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
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_ADDRESS")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	// dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/", password, host, dbport)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Fail to connect database")
	// }
	// db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	// if err != nil {
	// 	panic(err)
	// }

	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", password, host, dbport, dbname)
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
