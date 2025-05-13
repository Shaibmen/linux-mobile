package handlers

import (
	"errors"
	"os/exec"
	"server/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func ResourceMonitoring(c *gin.Context) {
	lines, err := RunAndSplit("free", "-h")
	if err != nil || len(lines) < 2 {
		c.JSON(500, err)
		return
	}

	memLines := strings.TrimSpace(lines[1])

	lines, err = RunAndSplit("df", "-h")
	if err != nil || len(lines) < 1 {
		c.JSON(500, err)
		return
	}

	diskLines := AddLines(lines)

	lines, err = RunAndSplit("top", "-b", "n1")
	if err != nil || len(lines) < 3 {
		c.JSON(500, err)
		return
	}

	cpuLines := []string{
		strings.TrimSpace(lines[0]),
		strings.TrimSpace(lines[0]),
		strings.TrimSpace(lines[0]),
	}

	c.JSON(200, models.Resource{
		Memory: memLines,
		Disk:   diskLines,
		CPU:    cpuLines,
	})

}

func NetworkMonitoring(c *gin.Context) {
	netstat, err := RunAndSplit("netstat", "-i")
	if err != nil {
		c.JSON(500, err)
		return
	}

	netLines := AddLines(netstat)

	ssi, err := RunAndSplit("ss", "-i")
	if err != nil {
		c.JSON(500, err)
		return
	}

	ssiLines := AddLines(ssi)

	c.JSON(200, models.Network{
		Netstat: netLines,
		Ssi:     ssiLines,
	})

}

func Split(out []byte) ([]string, error) {
	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		return nil, errors.New("неверный формат выхода команды")
	}
	return lines, nil
}

func RunAndSplit(cmd string, args ...string) ([]string, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return nil, err
	}

	lines, err := Split(out)
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func AddLines(lines []string) []string {

	var output []string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		output = append(output, line)
	}
	return output
}
