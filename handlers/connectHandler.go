package handlers

import (
	"net/http"
	"server/logging"
	"server/models"

	"github.com/gin-gonic/gin"
)

func ConnectPoint(c *gin.Context) {
	clientIP := c.ClientIP()
	c.JSON(http.StatusOK, models.HttpResponse{
		Out: "Успешное подключение к серверу!",
	})
	logging.Log.Info("Кто то подключился к машине: " + clientIP)
}
