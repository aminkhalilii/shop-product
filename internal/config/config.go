package config

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Postgres PostgresConfig
	Logger   LoggerConfig
}
