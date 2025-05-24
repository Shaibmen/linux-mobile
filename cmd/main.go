package main

import (
	"log"
	"server/environment"
	"server/handlers"
	"server/logging"
	"server/server"
)

func main() {

	err := logging.InitLogger("./app.log")
	if err != nil {
		log.Fatal("Ошибка инициализации логгера:", err)
	}
	responder := logging.InitJSONResponder()
	logging.ResponseJSON = responder
	logging.Log.Info("Http обработчик работает")

	environment.InitEnv()

	go handlers.CpuData()
	logging.Log.Info("Сбор инфомрации о процессоре начат")
	go handlers.RamData()
	logging.Log.Info("Сбор инфомрации о оперативной памяти начат")

	logging.Log.Info("Сервер запущен")
	server.StartServer()

}
