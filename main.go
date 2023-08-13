package main

import (
	"github.com/Ruberson1/api-gojobs/config"
	"github.com/Ruberson1/api-gojobs/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
