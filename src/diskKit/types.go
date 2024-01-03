package diskKit

type (
	// DiskUsageStats
	/*
		Total = Used + Free
	*/
	DiskUsageStats struct {
		//*disk.UsageStat

		Path string `json:"path" yaml:"path"`

		Free        uint64  `json:"free" yaml:"free"`
		Used        uint64  `json:"used" yaml:"used"`
		Total       uint64  `json:"total" yaml:"total"`
		UsedPercent float64 `json:"usedPercent" yaml:"usedPercent"`
	}
)
