package utils

import (
	"errors"
	"os/exec"
	"strings"
)

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
