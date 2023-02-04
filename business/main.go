package main

import (
	"business/initialize"
	"business/routes"

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
