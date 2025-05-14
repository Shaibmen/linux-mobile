package handlers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetFolder(c *gin.Context) {
	var request models.Folder
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	
}
