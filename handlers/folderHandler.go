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

	isdir, err := utils.CheckIsDir(request.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if isdir {
		folder, err := utils.RunAndSplit("ls", "-l", request.Path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		folderLines := utils.AddLines(folder)

		c.JSON(http.StatusOK, models.Folder{
			Files: folderLines,
		})
	} else {
		file, err := os.ReadFile(request.Path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, models.File{
			Data: string(file),
		})
	}

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
	isempty, err := utils.IsDirEmpty(request.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if !utils.IsSafePath(request.Path) {
		c.JSON(http.StatusBadRequest, gin.H{"out": "Запрещённый путь"})
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
		} else if isempty {
			os.Remove(request.Path)
			c.JSON(http.StatusOK, gin.H{"out": "Удаление завершено"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"out": "Папка не пуста, воспользуйтесь force"})
			return
		}
	}
}
