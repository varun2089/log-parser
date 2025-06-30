package main

import (
	"log-parser/internal/config"
	"os"
	"path/filepath"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

var testLogger = zerolog.Nop()

func TestRunApp_EmptyLogFilePath(t *testing.T) {
	cfg := config.Config{LogFilePath: ""}
	err := RunApp(cfg, testLogger)
	assert.Error(t, err)
	assert.EqualError(t, err, "log file path not set in the configuration")
}

func TestRunApp_ParseFail(t *testing.T) {
	cfg := config.Config{LogFilePath: "/nonexistent/file.log"}
	err := RunApp(cfg, testLogger)
	assert.Error(t, err)
}

func TestRunApp_Success(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.log")

	content := `
177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0"
168.41.191.40 - - [09/Jul/2018:10:11:30 +0200] "GET /faq/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0"
`

	err := os.WriteFile(filePath, []byte(content), 0644)
	assert.NoError(t, err)

	cfg := config.Config{LogFilePath: filePath}

	err = RunApp(cfg, testLogger)
	assert.NoError(t, err)
}
