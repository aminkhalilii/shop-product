package main

import (
	"log/slog"
	"product/internal/config"
	"product/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.New(cfg.Logger)

	log.Info("service started",
		slog.String("service", cfg.App.Name),
		slog.String("env", cfg.App.Env),
		slog.String("version", cfg.App.Version),
	)
}
