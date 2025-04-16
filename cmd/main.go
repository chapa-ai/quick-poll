package main

import (
	"fmt"
	"quick-poll/config"
	"quick-poll/internal/app"
	"quick-poll/pkg/logger"
)

func main() {
	log := logger.New()
	cfg, err := config.New()
	if err != nil {
		log.Error(fmt.Sprintf("config error: %s", err))
	}

	a := app.New(cfg, *log)
	if err = a.Run(); err != nil {
		log.Error(fmt.Sprintf("app run: %s", err))
	}
}
