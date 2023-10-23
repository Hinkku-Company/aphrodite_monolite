package logger

import (
	"os"

	"log/slog"
)

var logger *slog.Logger

func Log() *slog.Logger {
	return logger
}

func init() {
	var h slog.Handler
	h = slog.NewTextHandler(os.Stdout, nil)

	if value, exist := os.LookupEnv("APP_ENV"); exist && value == "production" {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})
	}

	logger = slog.New(h)
	slog.SetDefault(logger)
}
