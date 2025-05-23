package handlers

import (
	"net/http"
	"os"
	"server/logging"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func GetLogs(c *gin.Context) {

	logging.Log.Info("Получение логов")

	path := "./app.log"
	file, err := os.ReadFile(path)
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
