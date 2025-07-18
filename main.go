package main

import (
	"crud-oportunides/config"
	"crud-oportunides/router"
	"fmt"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	router.Initialize()
}