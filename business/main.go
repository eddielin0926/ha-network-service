package main

import (
	"ha-network-service/business/initialize"
	"ha-network-service/business/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnv()
}

func main() {
	app := gin.Default()
	routes.SetupRouter(app)
	app.Run()
}
