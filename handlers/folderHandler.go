package handlers

import (
	"net/http"
	"os"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func GetFolder(c *gin.Context) {
	var request models.Dir
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	exists, err := utils.CheckPath(request.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"out": "нет такого пути",
		})
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

func RemoveAny(c *gin.Context) {
	var request models.Dir
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	exists, err := utils.CheckPath(request.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"out": "нет такого пути",
		})
		return
	}

	isdir, err := utils.CheckIsDir(request.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	switch isdir {
	case false:
		os.Remove(request.Path)
		c.JSON(http.StatusOK, gin.H{"out": "Удаление завершено"})
	case true:
		if request.Force {
			os.RemoveAll(request.Path)
			c.JSON(http.StatusOK, gin.H{"out": "Форсированое удаление завершено"})
		} else {
			os.Remove(request.Path)
			c.JSON(http.StatusOK, gin.H{"out": "Удаление завершено"})
		}
	}

}
