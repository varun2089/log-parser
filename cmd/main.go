package main

import (
	"fmt"
	"log-parser/internal/config"
	"log-parser/internal/logger"
	"log-parser/internal/parser"
	"os"

	"github.com/rs/zerolog"
)

func RunApp(cfg config.Config, log zerolog.Logger) error {
	if cfg.LogFilePath == "" {
		log.Error().Msg("Log file path not set in the configuration")
		return fmt.Errorf("log file path not set in the configuration")
	}

	entries, err := parser.ParseLogFile(cfg.LogFilePath, log)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse log file")
		return err
	}

	uniqueIPs := parser.GetUniqueIPs(entries)
	topURLs := parser.GetTop3MostVisitedURLs(entries)
	activeIPs := parser.GetTop3MostActiveIPs(entries)

	log.Info().
		Int("unique_ips_count", len(uniqueIPs)).
		Msg("Unique IPs found in the log file")

	log.Info().
		Strs("top_3_urls", topURLs).
		Msg("Top 3 most visited URLs")

	log.Info().
		Strs("top_3_active_IPs", activeIPs).
		Msg("Top 3 active IPs")

	return nil
}

func main() {
	cfg := config.ParseConfig()
	log := logger.ConfigureLogger(cfg)

	if err := RunApp(cfg, log); err != nil {
		log.Error().Err(err).Msg("Application failed")
		os.Exit(1)
	}
}
