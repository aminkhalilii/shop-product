package config

import "github.com/joho/godotenv"

func Load() (*Config, error) {

	// در Production اگر فایل وجود نداشت نباید برنامه Fail شود.
	_ = godotenv.Load()

	cfg := &Config{

		App: AppConfig{
			Name:    getString("APP_NAME"),
			Env:     getString("APP_ENV"),
			Version: getString("APP_VERSION"),
		},

		Server: ServerConfig{
			Host:            getString("SERVER_HOST"),
			Port:            getInt("SERVER_PORT"),
			ReadTimeout:     getDuration("SERVER_READ_TIMEOUT"),
			WriteTimeout:    getDuration("SERVER_WRITE_TIMEOUT"),
			IdleTimeout:     getDuration("SERVER_IDLE_TIMEOUT"),
			ShutdownTimeout: getDuration("SERVER_SHUTDOWN_TIMEOUT"),
		},

		Postgres: PostgresConfig{
			Host:            getString("POSTGRES_HOST"),
			Port:            getInt("POSTGRES_PORT"),
			User:            getString("POSTGRES_USER"),
			Password:        getString("POSTGRES_PASSWORD"),
			Database:        getString("POSTGRES_DATABASE"),
			SSLMode:         getString("POSTGRES_SSLMODE"),
			MaxOpenConns:    getInt("POSTGRES_MAX_OPEN_CONNS"),
			MaxIdleConns:    getInt("POSTGRES_MAX_IDLE_CONNS"),
			ConnMaxLifetime: getDuration("POSTGRES_CONN_MAX_LIFETIME"),
			ConnMaxIdleTime: getDuration("POSTGRES_CONN_MAX_IDLE_TIME"),
		},

		Logger: LoggerConfig{
			Level: getString("LOGGER_LEVEL"),
			JSON:  getBool("LOGGER_JSON"),
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}
