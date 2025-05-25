package server

import (
	"server/handlers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	server := gin.Default()

	authorized := server.Group("/", middleware.CheckAuth())
	{
		authorized.GET("/enter", handlers.ConnectPoint)
		authorized.GET("/resource", handlers.ResourceMonitoring)
		authorized.GET("/network", handlers.NetworkMonitoring)

		authorized.GET("/process", handlers.GetProcess)
		authorized.POST("/process/grep", handlers.ProcessGrep)
		authorized.POST("/process/kill", handlers.Kill)
		authorized.POST("/process/terminate", handlers.Terminate)

		authorized.POST("/bash/create", handlers.CreateBash)
		authorized.POST("/bash/execute", handlers.ExecuteFile)
		authorized.GET("/bash/scripts", handlers.GetScripts)

		authorized.POST("/folder", handlers.GetFolder)
		authorized.POST("/folder/remove", handlers.RemoveAny)

		authorized.GET("/cpu/get", handlers.GetCPU)
		authorized.GET("/ram/get", handlers.GetRAM)

		authorized.GET("/logs", handlers.GetLogs)
	}

	server.Run(":3000")
}
