package server

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	server := gin.Default()

	server.GET("/enterPoint", handlers.ConnectPoint)
	server.GET("/resource", handlers.ResourceMonitoring)
	server.GET("/networkMonitoring", handlers.NetworkMonitoring)

	server.POST("/bashPoint_execute", handlers.ImportBashFile)

	server.Run(":8081")
}
