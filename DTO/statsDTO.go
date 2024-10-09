package DTO

type StatsResponseDTO struct {
	RamUsage              uint64 `json:"ramUsage"`
	RamUsageHumanReadable string `json:"ramUsageHumanReadable"`
	TotalRequests         uint64 `json:"totalRequests"`
	TotalDatabases        uint64 `json:"totalDatabases"`
	TotalKeys             uint64 `json:"totalKeys"`
	GoVersion             string `json:"goVersion"`
}
