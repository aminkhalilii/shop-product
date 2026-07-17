package config

import "time"

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}
