package environment

import (
	"os"
	"server/logging"

	"github.com/joho/godotenv"
)

type environment struct {
	AccessKey string
}

var Env environment

func InitEnv() {

	err := godotenv.Load()
	if err != nil {
		logging.Log.Error("Ошибка с загрузкой Env", err)
	}

	Env = environment{
		AccessKey: os.Getenv("API_KEY_ACCESS"),
	}
}
