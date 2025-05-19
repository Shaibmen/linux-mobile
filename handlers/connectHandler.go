package handlers

import (
	"net/http"
	"server/logging"

	"github.com/gin-gonic/gin"
)

func ConnectPoint(c *gin.Context) {
	clientIP := c.ClientIP()
	c.JSON(http.StatusOK, gin.H{"out": "Успешное подключение к серверу!"})
	logging.Log.Info("Кто то подключился к машине: " + clientIP)
}
