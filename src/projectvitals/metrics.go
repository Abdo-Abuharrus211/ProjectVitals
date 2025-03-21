package main

import(
	"time"
	"fmt"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/cpu"
)

func getCPUUsage() (float64, error) {
    cpuPercent, err := cpu.Percent(time.Second, false)
    if err != nil {
        fmt.Println("Failed to get CPU usage: ", err)
        return 0, err
    }
    return cpuPercent[0], nil
}

func getMemUsage() (float64, error) {
    virtMem, err := mem.VirtualMemory()
    if err != nil {
        fmt.Println("Failed to get memory usage: ", err)
        return 0, err
    }
    return virtMem.UsedPercent, nil
}


