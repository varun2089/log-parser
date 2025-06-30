package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTopNByKey_URLs(t *testing.T) {
	entries := []LogEntry{
		{URL: "/a"},
		{URL: "/b"},
		{URL: "/a"},
		{URL: "/c"},
		{URL: "/a"},
		{URL: "/b"},
		{URL: "/d"},
	}

	top := GetTopNByKey(entries, extractURL, 3)
	expected := []string{"/a", "/b", "/c"}

	assert.Equal(t, expected, top)
}

func TestGetTopNByKey_IPs(t *testing.T) {
	entries := []LogEntry{
		{IPAddress: "1.1.1.1"},
		{IPAddress: "2.2.2.2"},
		{IPAddress: "1.1.1.1"},
		{IPAddress: "3.3.3.3"},
		{IPAddress: "2.2.2.2"},
		{IPAddress: "1.1.1.1"},
		{IPAddress: "4.4.4.4"},
	}

	top := GetTopNByKey(entries, extractIP, 3)
	expected := []string{"1.1.1.1", "2.2.2.2", "3.3.3.3"}

	assert.Equal(t, expected, top)
}

func TestGetTopNByKey_Empty(t *testing.T) {
	var entries []LogEntry

	top := GetTopNByKey(entries, extractURL, 3)
	assert.Empty(t, top)
}
