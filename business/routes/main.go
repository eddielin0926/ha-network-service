package routes

import (
	"ha-network-service/business/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	api := app.Group("api")
	api.POST("/order", order)
	api.GET("/record", record)
	api.POST("/report", report)
}

func order(ctx *gin.Context) {
	var body models.Order
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Fail to read body"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func record(ctx *gin.Context) {
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

func report(ctx *gin.Context) {
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
