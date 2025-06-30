package logger

import (
	"github.com/rs/zerolog"
	"log-parser/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureLogger_ValidLevels(t *testing.T) {
	tests := []struct {
		name          string
		logLevel      string
		expectedLevel zerolog.Level
	}{
		{"debug level", "debug", zerolog.DebugLevel},
		{"info level", "info", zerolog.InfoLevel},
		{"warn level", "warn", zerolog.WarnLevel},
		{"error level", "error", zerolog.ErrorLevel},
		{"invalid level fallback", "invalid", zerolog.InfoLevel},
		{"case insensitive", "DeBuG", zerolog.DebugLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{LogLevel: tt.logLevel}
			_ = ConfigureLogger(cfg)

			assert.Equal(t, tt.expectedLevel, zerolog.GlobalLevel())
		})
	}
}

func TestSetLogLevel(t *testing.T) {
	setLogLevel(zerolog.ErrorLevel)
	assert.Equal(t, zerolog.ErrorLevel, zerolog.GlobalLevel())

	setLogLevel(zerolog.DebugLevel)
	assert.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel())
}
