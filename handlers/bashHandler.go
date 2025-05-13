package handlers

import (
	"fmt"
	"os/exec"
	structs "server/models"

	"github.com/gin-gonic/gin"
)

func ImportBashFile(c *gin.Context) {
	var bashfile structs.BashFile

	if err := c.BindJSON(&bashfile); err != nil {
		c.JSON(500, err)
		return
	}

	namefile := bashfile.NameField + ".bat"
	command := fmt.Sprintf("echo %s > %s", string(bashfile.TextField), namefile)
	cmd := exec.Command("cmd.exe", "/C", command)
	err := cmd.Run()
	if err != nil {
		c.JSON(500, nil)
		return
	}

	ExecuteFile(namefile)

	c.JSON(200, gin.H{"output": "200 ФАЙЛ СОЗДАН - ЗАПУЩЕН"})

}

func ExecuteFile(namefile string) {
	cmd := exec.Command("cmd.exe", "/C", namefile)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
