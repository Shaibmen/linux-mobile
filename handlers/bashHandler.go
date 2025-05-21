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

	const bashconst = "#!/bin/bash\n\n"

	filepath := homedir + "/bash_scripts/" + request.NameField + ".sh"

	if err := c.BindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверные данные", err)
		return
	}

	file, err := os.Create(filepath)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Ошибка создания bash", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(bashconst + request.TextField)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Ошибка записи bash", err)
		return
	}

	err = os.Chmod(filepath, 0755)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Не получается присвоить права", err)
		return
	}

	c.JSON(http.StatusOK, models.HttpResponse{
		Out: "Скрипт создан " + filepath,
	})
}

func ExecuteFile(c *gin.Context) {

	homedir := utils.HomeDir()

	var request structs.BashFile

	if err := c.BindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Неверные данные", err)
		return
	}

	cmd := exec.Command("bash", homedir+"/bash_scripts/"+request.NameField+".sh")
	output, err := cmd.CombinedOutput()
	result := models.BashOut{
		Out:   string(output),
		Error: " ",
	}
	if err != nil {
		result.Error = err.Error()
		utils.RespondWithError(c, http.StatusBadRequest, "Не получается выполнить скрипт", err)
		return
	}

	c.JSON(http.StatusOK, result)
}
