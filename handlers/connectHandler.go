package handlers

import "github.com/gin-gonic/gin"

func ConnectPoint(c *gin.Context) {
	c.JSON(200, gin.H{"code": "successful ping"})
}

