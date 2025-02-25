package handlers

import (
	"fmt"
	"os/exec"
	"server/structs"

	"github.com/gin-gonic/gin"
)

func ImportBashFile(c *gin.Context) {
	var bashfile structs.BashFile

	if err := c.BindJSON(&bashfile); err != nil {
		c.JSON(500, err)
		return
	}

	namefile := bashfile.NameField + ".bat"
	command := fmt.Sprintf("echo %v > %v", bashfile.TextField, namefile)
	cmd := exec.Command("cmd.exe", "/C", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"output": output})
}
