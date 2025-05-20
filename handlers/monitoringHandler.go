package handlers

import (
	"net/http"
	"server/models"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func ResourceMonitoring(c *gin.Context) {
	lines, err := utils.RunAndSplit("free", "-h")
	if err != nil || len(lines) < 2 {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удалось сформиовать ответ на ОЗУ", err)
		return
	}

	memLines := utils.AddLines(lines)

	lines, err = utils.RunAndSplit("df", "-h")
	if err != nil || len(lines) < 1 {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удалось сформиовать ответ на Storage", err)
		return
	}

	diskLines := utils.AddLines(lines)

	lines, err = utils.RunAndSplit("top", "-b", "n1")
	if err != nil || len(lines) < 3 {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удалось сформиовать ответ на CPU", err)
		return
	}

	cpuLines := []string{
		strings.TrimSpace(lines[0]),
		strings.TrimSpace(lines[1]),
		strings.TrimSpace(lines[2]),
	}

	c.JSON(http.StatusOK, models.Resource{
		Memory: memLines,
		Disk:   diskLines,
		CPU:    cpuLines,
	})

}

func NetworkMonitoring(c *gin.Context) {
	netstat, err := utils.RunAndSplit("netstat", "-i")
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удалось сформиовать ответ на NetStat", err)
		return
	}

	netLines := utils.AddLines(netstat)

	ssi, err := utils.RunAndSplit("ss", "-s")
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Не удалось сформиовать ответ на SSI", err)
		return
	}

	ssiLines := utils.AddLines(ssi)

	c.JSON(http.StatusOK, models.Network{
		Netstat: netLines,
		Ssi:     ssiLines,
	})

}
