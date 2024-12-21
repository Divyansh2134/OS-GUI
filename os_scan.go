package main

import (
	"fmt"
	"runtime"
	"strings"

	"example.com/m/models"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
)

func performScan() models.ScanResult {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Printf("Error fetching host info: %v\n", err)
		return models.ScanResult{}
	}

	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("Error fetching disk partitions: %v\n", err)
		return models.ScanResult{}
	}

	var disks []models.DiskInfo
	for _, partition := range partitions {
		if strings.HasPrefix(partition.Fstype, "sysfs") ||
			strings.HasPrefix(partition.Fstype, "proc") ||
			strings.HasPrefix(partition.Fstype, "cgroup") ||
			partition.Fstype == "" {
			continue
		}

		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Printf("Error fetching disk usage for %s: %v\n", partition.Device, err)
			continue
		}

		disks = append(disks, models.DiskInfo{
			Device:       partition.Device,
			FreeSpaceGB:  float64(usage.Free) / (1024 * 1024 * 1024),
			TotalSpaceGB: float64(usage.Total) / (1024 * 1024 * 1024),
		})
	}

	return models.ScanResult{
		OSName:      runtime.GOOS,
		OSVersion:   hostInfo.PlatformVersion,
		LastUpdate:  hostInfo.KernelVersion,
		DiskDetails: disks,
	}
}
