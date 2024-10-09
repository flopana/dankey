package HTTP

import (
	"dankey/DTO"
	"dankey/Util"
	"runtime"
)

func (s *Server) getStats() DTO.StatsResponseDTO {
	ramUsage := getMemoryUsage()
	res := DTO.StatsResponseDTO{
		RamUsage:              ramUsage,
		RamUsageHumanReadable: Util.ByteCountSI(int64(ramUsage)),
		TotalRequests:         s.requestCount.Load(),
		TotalDatabases:        s.Provider.GetTotalDatabases(),
		TotalKeys:             s.Provider.GetTotalKeys(),
		GoVersion:             runtime.Version(),
	}

	return res
}

func getMemoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m.Alloc
}
