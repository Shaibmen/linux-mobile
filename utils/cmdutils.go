package utils

import (
	"errors"
	"os"
	"os/exec"
	"os/user"
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

func IsSafePath(path string) bool {
	unsafe := []string{"/", "/bin", "/sbin", "/usr", "/lib", "/lib64", "/etc", "/var", "/boot", "/dev", "/proc", "/sys", "/root"}
	for _, p := range unsafe {
		if path == p || strings.HasPrefix(path, p+"/") {
			return false
		}
	}
	return true
}

func CheckPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckIsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func IsDirEmpty(path string) (bool, error) {
	infolder, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}

	return len(infolder) == 0, nil
}

func HomeDir() string {
	usr, _ := user.Current()
	homedir := usr.HomeDir
	return homedir
}
