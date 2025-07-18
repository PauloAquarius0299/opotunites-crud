package main

import (
	"crud-oportunides/config"
	"crud-oportunides/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//inicializando o config
	err := config.Init()
	if err != nil {
		logger.Error("Error initializing config:", err)
		return
	}
	router.Initialize()
}