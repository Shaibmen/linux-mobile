package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type environment struct {
	AccessKey string
}

var Env environment

func InitEnv() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	Env = environment{
		AccessKey: os.Getenv("API_KEY_ACCESS"),
	}
}
