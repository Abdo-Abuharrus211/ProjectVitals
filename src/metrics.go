package main

import (
	// "fmt"
	"os"
	// "runtime"
	"time"
    "strings"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	// "github.com/shirou/gopsutil/v4/process"
	"github.com/shirou/gopsutil/v4/sensors"
)

// Fetch PC name (hostname)
func GetPCName() (string, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostName, nil
}


// Fetch OS name and version
func GetOSInfo() (string, string, error) {
	info, err := host.Info()
	if err != nil {
		return "", "", err
	}
	return info.Platform, info.PlatformVersion, nil
}

// Fetch CPU usage percentage
func GetCPUStats() (string, float64, float64, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return "", 0.0, 0.0, err
	}
    cpuInfo, err := cpu.Info()
	if err != nil {
		return "", 0.0, 0.0, err
	}

    mainCPU := cpuInfo[0]
    cpuData := []string{mainCPU.Model, mainCPU.Model, mainCPU.CoreID}
    cpuName := strings.Join(cpuData, ",")

    var cpuTemp float64
    temps, err := sensors.SensorsTemperatures()
	if err != nil {
		return "", 0.0, 0.0, err
	}
	for _, temp := range temps {
		if temp.SensorKey == "coretemp" { // Adjust based on OS
			cpuTemp = temp.Temperature
		}
	}
	return cpuName, cpuPercent[0], cpuTemp, nil
}


// Fetch total memory and memory usage percentage
func GetMemoryStats() (uint64, float64, error) {
    virtMem, err := mem.VirtualMemory()
    if err != nil {
        return 0, 0, err
    }
    return virtMem.Total, virtMem.UsedPercent, nil
}

// Fetch disk usage
func GetDiskStats() (uint64, uint64, error) {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return 0, 0, err
	}
	return diskUsage.Total, diskUsage.Used, nil
}


// Fetch GPU details (limited support)
// func GetGPUStats() (string, float64, float64, error) {
// 	procs, err := process.Processes()
// 	if err != nil {
// 		return "", 0, 0, err
// 	}

// 	var gpuName string
// 	var gpuUsage float64
// 	var gpuTemp float64

// 	for _, proc := range procs {
// 		name, _ := proc.Name()
// 		if name == "nvidia-smi" { // Check for Nvidia GPUs (Linux)
// 			gpuName = "Nvidia GPU"
// 			gpuUsage = 50.0  // Fake data for now
// 			gpuTemp = 65.0   // Fake data
// 			break
// 		}
//         name, _ := proc.Name()
//         if name == "radeontop" { // Check for AMD GPUs (Linux)
//             gpuName = "AMD GPU"
//             gpuUsage = 50.0  // Fake data
//             gpuTemp = 65.0   // Fake data
//             break
//         }
// 	}

// 	if gpuName == "" {
// 		return "Unknown GPU", 0, 0, fmt.Errorf("GPU information not available")
// 	}

// 	return gpuName, gpuUsage, gpuTemp, nil
// }
