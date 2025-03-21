package main

import (
	"fmt"
	"math"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
)

const refreshInterval = time.Second // Update interval

// Model struct to store app state
type model struct {
	cursor     int
	spinner    SpinnerModel
	PCName     string
	OSName     string
	OSVersion  string
	CPUName    string
	CPUPercent float64
	CPUTemp    float64
	MemTotal   uint64
	MemPercent float64
	DiskTotal  uint64
	DiskUsed   uint64
	GPUName    string
	GPUPercent float64
	GPUTemp    float64
}

// Initializes the model
func initialModel() model {
	spin := spinner.New()
	spin.Spinner = spinner.Pulse // Default spinner type
	spin.Style = spinnerStyle

	return model{
		cursor:  0,
		spinner: SpinnerModel{Spinner: spin},
	}
}

// Init function runs on program start
func (m model) Init() tea.Cmd {
	return tea.Batch(refresh(), m.spinner.Spinner.Tick)
}

// refreshMsg struct to trigger periodic updates
type refreshMsg struct{}

// Function to trigger periodic updates
func refresh() tea.Cmd {
	return tea.Tick(refreshInterval, func(time.Time) tea.Msg {
		return refreshMsg{}
	})
}

// Update function handles messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case refreshMsg:
		PCName, _ := GetPCName()
		osName, osVersion, _ := GetOSInfo()
		cpuName, cpuUse, cpuTemp, _ := GetCPUStats()
		memTotal, memUse, _ := GetMemoryStats()
		diskTotal, diskUsed, _ := GetDiskStats()

		m.PCName = PCName
		m.OSName, m.OSVersion = osName, osVersion
		m.CPUName, m.CPUPercent, m.CPUTemp = cpuName, cpuUse, cpuTemp
		m.MemTotal = memTotal
		m.MemPercent = memUse
		m.DiskTotal = diskTotal
		m.DiskUsed = diskUsed

		return m, refresh() // Schedule next refresh

	case spinner.TickMsg:
		var spinCmd tea.Cmd
		m.spinner.Spinner, spinCmd = m.spinner.Spinner.Update(msg)
		return m, spinCmd
	}

	return m, cmd
}

/*
	Render UI in terminal.
	Parses the model's state and stringies it, the string becomes the UI.
	BubbleTea handles all the redrawing.
*/
func (m model) View() string {
	// TODO: add the logic for how we render stuff.
    s := "Project Vitals\n\n" //TODO: update name
    s += fmt.Sprintf("Monitoring vitals %s\n", m.spinner.Spinner.View())
	s += fmt.Sprintf("PC Name: '%s'\n", m.PCName)
	s += fmt.Sprintf("OS: %s %s\n", m.OSName, m.OSVersion)
	s += fmt.Sprintf("CPU: %s\n", m.CPUName)
	s += fmt.Sprintf("CPU Usage: %.1f%%\n", m.CPUPercent)
	s += fmt.Sprintf("CPU Temperature: %.1f°C\n", m.CPUTemp)
	s += fmt.Sprintf("Memory Total: %d bits (%.2f GB)\n", m.MemTotal, float64(m.MemTotal) / math.Pow10(9))
	s += fmt.Sprintf("Memory Usage: %.1f%%\n", m.MemPercent)
	s += fmt.Sprintf("Disk Total: %.2f GB\n", float64(m.DiskTotal) / math.Pow10(9))
    diskUsagePercent := 0.0
    if m.DiskTotal > 0 {
        diskUsagePercent = (float64(m.DiskUsed) / float64(m.DiskTotal)) * 100
    }
    s += fmt.Sprintf("Disk Usage: %.1f%%\n", diskUsagePercent)
    s += RenderProgressBar(diskUsagePercent, 30) + "\n"
	s += fmt.Sprintf("Disk Free: %.2f GB\n", float64(m.DiskTotal - m.DiskUsed) / math.Pow10(9))
	// Uncomment the following lines if GPU stats become available
	// s += fmt.Sprintf("GPU: %s\n", m.GPUName)
	// s += fmt.Sprintf("GPU Usage: %.1f%%\n", m.GPUPercent)
	// s += fmt.Sprintf("GPU Temperature: %.1f°C\n", m.GPUTemp)
    s += "\nPress q to quit\n"
    return s
}
/*
	Main function to run the program.
*/
func main(){
	monitorProgram := tea.NewProgram(initialModel())
	if _, err := monitorProgram.Run(); err != nil {
		fmt.Printf("Error occured: %v", err)
		os.Exit(1)
	}
}
