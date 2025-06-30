package parser

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"regexp"
)

var logLinePattern = regexp.MustCompile(`^(\d+\.\d+\.\d+\.\d+).+"[A-Z]+ (.+?) HTTP/`)

func ParseLogFile(logfilePath string, log zerolog.Logger) ([]LogEntry, error) {

	file, err := os.Open(logfilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open log file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Warn().Err(err).Msg("error closing log file")
		}
	}()

	var entries []LogEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := logLinePattern.FindStringSubmatch(line)
		if len(matches) == 3 {
			entry := LogEntry{
				IPAddress: matches[1],
				URL:       matches[2],
			}
			entries = append(entries, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading log file: %w", err)
	}

	return entries, nil
}
