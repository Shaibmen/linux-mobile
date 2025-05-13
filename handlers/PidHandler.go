package handlers

import (
	"net/http"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func ProccesGet(c *gin.Context) {

	process, err := utils.RunAndSplit("top")
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

	ps, err := utils.RunAndSplit("ps", "aux", "|", "grep", request.Prefix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	psLines := utils.AddLines(ps)

	c.JSON(http.StatusOK, models.Process{
		Process: psLines,
	})
}
