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
		utils.RespondWithError(c, http.StatusBadRequest, "Неправильные данные", err)
		return
	}

	exists, err := utils.CheckPath(request.Path)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Ошибка проверки пути", err)
		return
	}
	if !exists {
		utils.RespondWithError(c, http.StatusBadRequest, "Нет такого пути", err)
		return
	}

	isdir, err := utils.CheckIsDir(request.Path)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Ошибка проверки на папку", err)
		return
	}

	if isdir {
		folder, err := utils.RunAndSplit("ls", "-l", request.Path)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Не удаётся сформировать ответ", err)
			return
		}

		folderLines := utils.AddLines(folder)

		c.JSON(http.StatusOK, models.Folder{
			Files: folderLines,
		})
	} else {
		file, err := os.ReadFile(request.Path)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Не удаётся прочитать файл", err)
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
		utils.RespondWithError(c, http.StatusBadRequest, "Неправильные данные", err)
		return
	}
	exists, err := utils.CheckPath(request.Path)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удаётся проверить путь", err)
		return
	}
	if !exists {
		utils.RespondWithError(c, http.StatusBadRequest, "Такого пути не существует", err)
		return
	}

	isdir, err := utils.CheckIsDir(request.Path)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Ошибка проверки на папку", err)
		return
	}
	isempty, err := utils.IsDirEmpty(request.Path)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Проверка на вложенность папки не удалась", err)
		return
	}

	if !utils.IsSafePath(request.Path) {
		utils.RespondWithError(c, http.StatusInternalServerError, "Запрещённый путь", err)
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
			utils.RespondWithError(c, http.StatusBadRequest, "Папка не пуста воспользуйтесь флагом Force", err)
			return
		}
	}
}
