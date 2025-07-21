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
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	router.Initialize()
}