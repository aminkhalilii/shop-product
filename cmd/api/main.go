package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"strconv"

	"product/internal/config"
	"product/internal/database"
	"product/internal/handler"
	"product/internal/logger"
	"product/internal/repository/postgres"
	"product/internal/router"
	"product/internal/service"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.New(cfg.Logger)

	ctx := context.Background()

	db, err := database.NewPostgres(
		ctx,
		cfg.Postgres,
	)

	if err != nil {
		log.Error(
			"database connection failed",
			slog.Any("error", err),
		)

		os.Exit(1)
	}

	defer db.Close()

	log.Info("database connected")

	log.Info(
		"service started",
		slog.String("service", cfg.App.Name),
		slog.String("env", cfg.App.Env),
		slog.String("version", cfg.App.Version),
	)

	productRepo := postgres.NewProductRepository(db)
	productService := service.NewProductService(
		productRepo,
	)
	productHandler := handler.NewProductHandler(
		productService,
	)
	r := router.NewRouter(
		productHandler,
	)
	addr := net.JoinHostPort(
		cfg.Server.Host,
		strconv.Itoa(cfg.Server.Port),
	)
	err = r.Run(addr)

	if err != nil {
		panic(err)
	}

}
