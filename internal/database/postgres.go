package database

import (
	"context"
	"fmt"
	"time"

	"product/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(
	ctx context.Context,
	cfg config.PostgresConfig,
) (*pgxpool.Pool, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse postgres config: %w", err)
	}

	// Connection Pool Settings

	poolConfig.MaxConns = int32(cfg.MaxOpenConns)

	poolConfig.MinConns = int32(cfg.MaxIdleConns)

	poolConfig.MaxConnLifetime = cfg.ConnMaxLifetime

	poolConfig.MaxConnIdleTime = cfg.ConnMaxIdleTime

	pool, err := pgxpool.NewWithConfig(
		ctx,
		poolConfig,
	)

	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	// Check Database Connection

	pingCtx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)

	defer cancel()

	err = pool.Ping(pingCtx)

	if err != nil {
		pool.Close()

		return nil, fmt.Errorf("postgres ping failed: %w", err)
	}

	return pool, nil
}
