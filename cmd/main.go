package main

import (
	"server/environment"
	"server/server"
)

func main() {

	environment.InitEnv()
	server.StartServer()
}
