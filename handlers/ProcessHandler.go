package handlers

import (
	"net/http"
	"server/models"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func ProccesGet(c *gin.Context) {

	process, err := utils.RunAndSplit("top", "-b", "-n", "1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	ps, err := utils.RunAndSplit("ps", "aux")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
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
