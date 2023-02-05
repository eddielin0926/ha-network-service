package routes

import (
	"context"
	"ha-network-service/business/models"
	"ha-network-service/grpcpb/inventory"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiRoute interface {
	Setup(app *gin.Engine)
}

type apiRoute struct {
	inventoryClient inventory.InventoryClient
}

func NewApiRoute(c inventory.InventoryClient) ApiRoute {
	return &apiRoute{
		inventoryClient: c,
	}
}

func (r *apiRoute) Setup(app *gin.Engine) {
	api := app.Group("api")
	api.POST("/order", r.order)
	api.GET("/record", r.record)
	api.POST("/report", r.report)
}

func (r *apiRoute) order(c *gin.Context) {
	var body models.Order
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail to read body"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ret, err := r.inventoryClient.GetMaterial(ctx, &inventory.Order{
		A: int32(body.Data.A),
		B: int32(body.Data.B),
		C: int32(body.Data.C),
		D: int32(body.Data.D),
	})
	if err != nil {
		log.Fatalf("could not get material: %v", err)
	}
	log.Printf("Material: %v, Signature: %v", ret.GetMaterial(), ret.GetSignature())
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (r *apiRoute) record(ctx *gin.Context) {
	var body struct {
		Location string    `json:"location"`
		Date     time.Time `json:"date"`
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Fail to read body"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (r *apiRoute) report(ctx *gin.Context) {
	var body struct {
		Location string    `json:"location"`
		Date     time.Time `json:"date"`
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Fail to read body"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
