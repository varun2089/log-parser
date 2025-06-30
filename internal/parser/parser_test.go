package parser

import (
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestParseLogFile_Success(t *testing.T) {
	content := `
177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0"
168.41.191.40 - - [09/Jul/2018:10:11:30 +0200] "GET /faq/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0"
invalid log line without IP
`
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.log")
	err := os.WriteFile(filePath, []byte(content), 0644)
	assert.NoError(t, err)

	log := zerolog.Nop()

	entries, err := ParseLogFile(filePath, log)
	assert.NoError(t, err)

	expected := []LogEntry{
		{IPAddress: "177.71.128.21", URL: "/intranet-analytics/"},
		{IPAddress: "168.41.191.40", URL: "/faq/"},
	}

	assert.Equal(t, expected, entries)
}

func TestParseLogFile_FileNotFound(t *testing.T) {
	log := zerolog.Nop()
	entries, err := ParseLogFile("nonexistent.log", log)
	assert.Nil(t, entries)
	assert.Error(t, err)
}

func TestParseLogFile_EmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	emptyFile := filepath.Join(tmpDir, "empty.log")
	err := os.WriteFile(emptyFile, []byte(""), 0644)
	assert.NoError(t, err)

	log := zerolog.Nop()
	entries, err := ParseLogFile(emptyFile, log)
	assert.NoError(t, err)
	assert.Empty(t, entries)
}
