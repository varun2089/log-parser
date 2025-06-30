package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUniqueIPs(t *testing.T) {
	entries := []LogEntry{
		{IPAddress: "1.1.1.1"},
		{IPAddress: "2.2.2.2"},
		{IPAddress: "1.1.1.1"},
		{IPAddress: "3.3.3.3"},
	}

	unique := GetUniqueIPs(entries)

	expected := map[string]struct{}{
		"1.1.1.1": {},
		"2.2.2.2": {},
		"3.3.3.3": {},
	}

	assert.Equal(t, expected, unique)
}

func TestGetTop3MostVisitedURLs(t *testing.T) {
	entries := []LogEntry{
		{URL: "/a"},
		{URL: "/b"},
		{URL: "/a"},
		{URL: "/c"},
		{URL: "/a"},
		{URL: "/b"},
		{URL: "/d"},
	}

	top := GetTop3MostVisitedURLs(entries)

	expected := []string{"/a", "/b", "/c"}

	assert.Equal(t, expected, top)
}

func TestGetTop3MostActiveIPs(t *testing.T) {
	entries := []LogEntry{
		{IPAddress: "1.1.1.1"},
		{IPAddress: "2.2.2.2"},
		{IPAddress: "1.1.1.1"},
		{IPAddress: "3.3.3.3"},
		{IPAddress: "2.2.2.2"},
		{IPAddress: "1.1.1.1"},
		{IPAddress: "4.4.4.4"},
	}

	top := GetTop3MostActiveIPs(entries)

	expected := []string{"1.1.1.1", "2.2.2.2", "3.3.3.3"}

	assert.Equal(t, expected, top)
}
