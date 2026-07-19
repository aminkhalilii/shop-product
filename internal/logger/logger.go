package logger

import (
	"log/slog"
	"os"

	"product/internal/config"
)

func New(cfg config.LoggerConfig) *slog.Logger {

	var handler slog.Handler

	if cfg.JSON {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: parseLevel(cfg.Level),
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: parseLevel(cfg.Level),
		})
	}

	return slog.New(handler)
}

func parseLevel(level string) slog.Level {

	switch level {

	case "debug":
		return slog.LevelDebug

	case "warn":
		return slog.LevelWarn

	case "error":
		return slog.LevelError

	default:
		return slog.LevelInfo
	}
}
