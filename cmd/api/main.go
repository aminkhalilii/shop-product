package main

import (
	"product/internal/config"
	"product/internal/logger"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}

	log := logger.New(cfg.Logger)

	log.Info(
		"product service started",
		"name", cfg.App.Name,
	)
}
