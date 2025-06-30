package logger

import (
	"github.com/rs/zerolog"
	"log-parser/internal/config"
	"os"
	"strings"
)

func DefaultLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func setLogLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}

func ConfigureLogger(cfg config.Config) zerolog.Logger {
	log := DefaultLogger()

	level, err := zerolog.ParseLevel(strings.ToLower(cfg.LogLevel))
	if err != nil {
		level = zerolog.InfoLevel // fallback to Info if invalid level
	}

	setLogLevel(level)

	return log
}
