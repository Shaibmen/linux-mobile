package server

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	server := gin.Default()

	server.GET("/enter", handlers.ConnectPoint)
	server.GET("/resource", handlers.ResourceMonitoring)
	server.GET("/network", handlers.NetworkMonitoring)

	server.GET("/process", handlers.GetProcess)
	server.POST("/process/grep", handlers.ProcessGrep)
	server.POST("/process/kill", handlers.Kill)
	server.POST("/process/terminate", handlers.Terminate)

	server.POST("/bash/execute", handlers.ImportBashFile)

	server.POST("/folder", handlers.GetFolder)

	server.Run(":3000")
}
