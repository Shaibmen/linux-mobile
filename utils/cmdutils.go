package utils

import (
	"errors"
	"os"
	"os/exec"
	"os/user"

	"server/logging"
	"strings"

	"github.com/gin-gonic/gin"
)

func Split(out []byte) ([]string, error) {
	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		logging.Log.Error("Ожидаемый выход меньше чем 2 строки", nil)
		return nil, errors.New("неверный формат выхода команды")
	}
	return lines, nil
}

func RunAndSplit(cmd string, args ...string) ([]string, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		logging.Log.Error("Ошибка в RunAndSplit", err)
		return nil, err
	}

	lines, err := Split(out)
	if err != nil {
		logging.Log.Error("Ошибка в RunAndSplit - вызов Split", err)
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
		logging.Log.Error("Ошибка в CheckPath", err)
		return false, err
	}
	return true, nil
}

func CheckIsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		logging.Log.Error("Ошибка в CheckIsDir", err)
		return false, err
	}
	return info.IsDir(), nil
}

func IsDirEmpty(path string) (bool, error) {
	infolder, err := os.ReadDir(path)
	if err != nil {
		logging.Log.Error("Ошибка в IsDirEmpty", err)
		return false, err
	}

	return len(infolder) == 0, nil
}

func HomeDir() string {
	usr, _ := user.Current()
	homedir := usr.HomeDir
	return homedir
}

func RespondWithError(c *gin.Context, status int, msg string, err error) {
	logging.ResponseJSON.Error(c, status, msg)
	logging.Log.Error(msg, err)
}
