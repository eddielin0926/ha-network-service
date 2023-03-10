package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/eddielin0926/ha-network-service/business/models"
	"github.com/eddielin0926/ha-network-service/grpcpb"
	"github.com/eddielin0926/ha-network-service/grpcpb/inventory"
	"github.com/eddielin0926/ha-network-service/grpcpb/storage"

	"github.com/gin-gonic/gin"
)

type ApiRoute interface {
	Setup(app *gin.Engine)
}

type apiRoute struct {
	inventoryClient inventory.InventoryClient
	storageClient   storage.StorageClient
}

func NewApiRoute(invc inventory.InventoryClient, storc storage.StorageClient) ApiRoute {
	return &apiRoute{
		inventoryClient: invc,
		storageClient:   storc,
	}
}

func (r *apiRoute) Setup(app *gin.Engine) {
	api := app.Group("api")
	api.POST("/order", r.order)
	api.GET("/record", r.record)
	api.GET("/report", r.report)
}

func (r *apiRoute) order(c *gin.Context) {
	var body models.Order
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail to read body"})
		return
	}

	invCtx, invCancel := context.WithTimeout(context.Background(), time.Second)
	defer invCancel()
	invRet, err := r.inventoryClient.GetMaterial(invCtx, &inventory.Order{
		A: int32(body.Data.A),
		B: int32(body.Data.B),
		C: int32(body.Data.C),
		D: int32(body.Data.D),
	})
	if err != nil {
		log.Printf("could not get material: %v", err)
		return
	}

	storCtx, storCancel := context.WithTimeout(context.Background(), time.Second)
	defer storCancel()
	storRet, err := r.storageClient.SaveRecord(storCtx, &storage.Record{
		Location:  body.Location,
		Timestamp: body.Timestamp,
		Signature: invRet.Signature,
		Material:  invRet.Material,
		A:         body.Data.A,
		B:         body.Data.B,
		C:         body.Data.C,
		D:         body.Data.D,
	})
	if err != nil {
		log.Printf("could not store record: %v", err)
		return
	}
	if storRet.Status == storage.Status_FAIL {
		log.Printf("fail to store record")
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (r *apiRoute) record(ctx *gin.Context) {
	location := ctx.Query("location")
	if location == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing location"})
		return
	}
	date := ctx.Query("date")
	if date == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing date"})
		return
	}

	storCtx, storCancel := context.WithTimeout(context.Background(), time.Second)
	defer storCancel()
	storRet, err := r.storageClient.GetRecords(storCtx, &storage.Query{
		Location: location,
		Date:     date,
	})
	if err != nil {
		log.Printf("could not store record: %v", err)
		return
	}

	bin, _ := grpcpb.ProtobufToJson(storRet)
	var data models.Records
	err = json.Unmarshal(bin, &data)
	if err != nil {
		log.Println("Unmarshal fail")
	}

	ctx.JSON(http.StatusOK, data.Records)
}

func (r *apiRoute) report(ctx *gin.Context) {
	location := ctx.Query("location")
	if location == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing location"})
		return
	}
	date := ctx.Query("date")
	if date == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing date"})
		return
	}

	storCtx, storCancel := context.WithTimeout(context.Background(), time.Second)
	defer storCancel()
	storRet, err := r.storageClient.GetReport(storCtx, &storage.Query{
		Location: location,
		Date:     date,
	})
	if err != nil {
		log.Printf("could not store record: %v", err)
		return
	}

	bin, _ := grpcpb.ProtobufToJson(storRet)

	ctx.Data(http.StatusOK, gin.MIMEJSON, bin)
}
