package models

type DiskInfo struct {
	Device       string  `json:"device"`
	FreeSpaceGB  float64 `json:"free_space_gb"`
	TotalSpaceGB float64 `json:"total_space_gb"`
}
