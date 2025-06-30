package parser

func GetUniqueIPs(entries []LogEntry) map[string]struct{} {
	uniqueIPs := make(map[string]struct{})
	for _, entry := range entries {
		uniqueIPs[entry.IPAddress] = struct{}{}
	}
	return uniqueIPs
}

func GetTop3MostVisitedURLs(entries []LogEntry) []string {
	return GetTopNByKey(entries, extractURL, 3)
}

func GetTop3MostActiveIPs(entries []LogEntry) []string {
	return GetTopNByKey(entries, extractIP, 3)
}
