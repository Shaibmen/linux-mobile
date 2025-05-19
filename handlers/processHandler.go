package handlers

import (
	"net/http"
	"os/exec"
	"server/models"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetProcess(c *gin.Context) {

	process, err := utils.RunAndSplit("top", "-b", "-n", "1")
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удалось сформиовать ответ на top", err)
		return
	}
	processLines := utils.AddLines(process)

	c.JSON(http.StatusOK, models.Process{
		Process: processLines,
	})
}

func ProcessGrep(c *gin.Context) {

	var request models.ProcessGrep
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "не правильные данные", err)
		return
	}

	ps, err := utils.RunAndSplit("ps", "aux")
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удалось сформиовать ответ на ps aux", err)
		return
	}

	var psLines []string
	for _, line := range ps {
		if strings.Contains(line, request.Prefix) {
			psLines = append(psLines, line)
		}
	}

	c.JSON(http.StatusOK, models.Process{
		Process: psLines,
	})
}

func Kill(c *gin.Context) {
	var request models.PID
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "не правильные данные", err)
		return
	}

	output, err := exec.Command("kill", request.ID).Output()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удалось мягко завершить процесс", err)
		return
	}

	c.JSON(http.StatusAccepted, output)
}

func Terminate(c *gin.Context) {
	var request models.PID
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "не правильные данные", err)
		return
	}

	output, err := exec.Command("kill", "-9", request.ID).Output()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удалось жёстко завершить процесс", err)
		return
	}
	c.JSON(http.StatusAccepted, output)
}
