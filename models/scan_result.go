package models

type ScanResult struct {
	OSName      string `json:"os_name"`
	OSVersion   string `json:"os_version"`
	LastUpdate  string `json:"last_update"`
	DiskDetails []DiskInfo
}
