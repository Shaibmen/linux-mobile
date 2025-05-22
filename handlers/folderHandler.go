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
			Data: folderLines,
		})
	} else {
		file, err := os.ReadFile(request.Path)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Не удаётся прочитать файл", err)
			return
		}
		lines, err := utils.Split(file)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Не удаётся запарсить файл", err)
			return
		}
		c.JSON(http.StatusOK, models.File{
			Data: lines,
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

	if !utils.IsSafePath(request.Path) && request.IgnoreAlert == false {
		utils.RespondWithError(c, http.StatusConflict, "Взаимодействие в этой директории может вызвать крах системы. \n Действуйте на свой страх и риск.", err)
		return
	}

	switch isdir {
	case false:
		os.Remove(request.Path)
		c.JSON(http.StatusOK, models.HttpResponse{
			Out: "Удаление завершено",
		})
	case true:
		isempty, err := utils.IsDirEmpty(request.Path)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Проверка на вложенность папки не удалась", err)
			return
		}
		if request.Force {
			os.RemoveAll(request.Path)
			c.JSON(http.StatusOK, models.HttpResponse{
				Out: "Форсированое удаление завершено",
			})
		} else if isempty {
			os.Remove(request.Path)
			c.JSON(http.StatusOK, models.HttpResponse{
				Out: "Удаление завершено",
			})
		} else {
			utils.RespondWithError(c, http.StatusBadRequest, "Папка не пуста воспользуйтесь флагом Force", err)
			return
		}
	}
}
