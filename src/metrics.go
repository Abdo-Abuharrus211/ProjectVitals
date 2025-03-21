package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

// Fetch CPU usage percentage
func GetCPUUsage() (float64, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Failed to get CPU usage:", err)
		return 0, err
	}
	return cpuPercent[0], nil
}

// Fetch memory usage percentage
func GetMemUsage() (float64, error) {
	virtMem, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Failed to get memory usage:", err)
		return 0, err
	}
	return virtMem.UsedPercent, nil
}
