package config

import (
	"github.com/alecthomas/kong"
)

type Config struct {
	LogLevel    string `env:"LOG_LEVEL" default:"INFO" help:"Log level for the application"`
	LogFilePath string `env:"LOG_FILE_PATH" default:"access.log" help:"Path to the log file to be parsed"`
}

func ParseConfig() Config {
	var config Config
	kong.Parse(&config)
	return config
}
