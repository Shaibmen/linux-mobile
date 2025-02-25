package handlers

import (
	"os/exec"

	"github.com/gin-gonic/gin"
)

func ResourceMonitoring(c *gin.Context) {
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && dir")
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"output": string(output)})

}

func NetworkMonitoring(c *gin.Context) {
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && ipconfig")
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, gin.H{"output": string(output)})
}
