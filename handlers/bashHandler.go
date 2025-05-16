package handlers

import (
	"net/http"
	"os"
	"os/exec"
	"server/models"
	structs "server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func CreateBash(c *gin.Context) {

	homedir := utils.HomeDir()

	var request structs.BashFile

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	file, err := os.Create(homedir + "/bash_scripts/" + request.NameField + ".sh")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	_, err = file.WriteString(request.TextField)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = os.Chmod(homedir+"/bash_scripts/"+request.NameField+".sh", 0755)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
}

func ExecuteFile(c *gin.Context) {

	homedir := utils.HomeDir()

	var request structs.BashFile

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	cmd := exec.Command("bash", homedir+"/bash_scripts/"+request.NameField+".sh")
	output, err := cmd.CombinedOutput()
	result := models.BashOut{
		Out:   string(output),
		Error: "",
	}
	if err != nil {
		result.Error = err.Error()
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
