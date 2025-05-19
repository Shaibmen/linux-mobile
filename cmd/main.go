package main

import (
	"log"
	"server/environment"
	"server/logging"
	"server/server"
)

var Log *logging.FileLogger
var err error

func main() {

	Log, err = logging.NewFileLogger("./app.log")
	if err != nil {
		log.Fatal("Ошибка инициализации логгера")
	}
	environment.InitEnv()
	server.StartServer()

	Log.Info("Сервер запущен")
}
