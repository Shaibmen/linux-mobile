package handlers

import (
	"net/http"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func GetFolder(c *gin.Context) {
	var request models.FileJson
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	folder, err := utils.RunAndSplit("ls", "-l", request.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	folderLines := utils.AddLines(folder)

	c.JSON(http.StatusOK, models.Folder{
		Files: folderLines,
	})
}
