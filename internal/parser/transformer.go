package parser

import "sort"

func GetTopNByKey(entries []LogEntry, keyFunc func(LogEntry) string, n int) []string {
	counts := make(map[string]int)
	for _, e := range entries {
		key := keyFunc(e)
		counts[key]++
	}

	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return counts[keys[i]] > counts[keys[j]]
	})

	if len(keys) > n {
		return keys[:n]
	}
	return keys
}

func extractURL(e LogEntry) string {
	return e.URL
}

func extractIP(e LogEntry) string {
	return e.IPAddress
}
