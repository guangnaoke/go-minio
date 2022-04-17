package main

import (
	"minio_server/initialize"
	"minio_server/router"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a Minio Server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @securityDefinitions.apikey ApiKeyXRole
// @in header
// @name x-role
// @host localhost:8082
// @BasePath /
func main() {
	color.Yellow("mysql & minio =====> Init.....")
	initialize.Init()
	color.Yellow("Init end!")

	color.Red("=========")

	color.Blue("gin service started ======>")
	r := gin.Default()
	r = router.CollectRoute(r)
	r.Run(":8082")
}
