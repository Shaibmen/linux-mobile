package handlers

import (
	"errors"
	"os/exec"
	"server/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func ResourceMonitoring(c *gin.Context) {
	out, err := exec.Command("free", "-h").Output()
	if err != nil {
		c.JSON(500, err)
	}

	lines, err := Split(out)
	if err != nil {
		c.JSON(500, err)
	}

	memLine := strings.Fields(lines[1])

	out, err = exec.Command("df", "-h").Output()
	if err != nil {
		c.JSON(500, err)
	}

	lines, err = Split(out)
	if err != nil {
		c.JSON(500, err)
	}

	diskLine := strings.Fields(lines[1])

	cpuLine, err := exec.Command("ip", "-s", "link").Output()
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, models.Resource{
		Memory: memLine,
		Disk:   diskLine,
		CPU:    string(cpuLine),
	})

}

func NetworkMonitoring(c *gin.Context) {
	cmd := exec.Command("cmd.exe", "/C", "chcp 65001 && ipconfig")
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, gin.H{"output": string(output)})
}

func Split(out []byte) ([]string, error) {
	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		return nil, errors.New("неверный формат вывода free -h")
	}
	return lines, nil
}
