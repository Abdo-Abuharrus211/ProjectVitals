package main

import (
	"fmt"
	"math"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
    progress "github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
)

const refreshInterval = time.Second // Update interval

// Model struct to store app state
type model struct {
	cursor     int
	spinner    SpinnerModel
    progressBar progress.Model

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
    pbar := progress.New(progress.WithDefaultGradient())
	spin := spinner.New()
	spin.Spinner = spinner.Line // Default spinner type
	spin.Style = spinnerStyle

	return model{
        cursor:      0,
        spinner:     SpinnerModel{Spinner: spin},
        progressBar: pbar,
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

        diskUsage := 0.0
        if m.DiskTotal > 0 {
            diskUsage = float64(m.DiskUsed) / float64(m.DiskTotal)
        }
        progressCmd := m.progressBar.SetPercent(diskUsage)
        return m, tea.Batch(refresh(), progressCmd)
    case progress.FrameMsg:
        progressModel, progressCmd := m.progressBar.Update(msg)
        m.progressBar = progressModel.(progress.Model)
        return m, progressCmd
        
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
    // s := "Project Vitals\n\n" //TODO: update name
    s := fmt.Sprintf("%s\n\n", titleStyle("ProjectVitalis"))
    s += fmt.Sprintf("%s%s\n", labelStyle("Monitoring system vitals "), m.spinner.Spinner.View())
    s += fmt.Sprintf("%s%s\n", labelStyle("PC Name: "), textStyle("'" + m.PCName + "'"))
    s += fmt.Sprintf("%s%s%s\n", labelStyle("OS: "), textStyle(m.OSName), textStyle(m.OSVersion))
    s += fmt.Sprintf("%s%s\n", labelStyle("CPU: "), textStyle(m.CPUName))
    s += fmt.Sprintf("%s%s\n", labelStyle("CPU Usage: "), textStyle(fmt.Sprintf("%.1f%%", m.CPUPercent)))
    s += fmt.Sprintf("%s%s\n", labelStyle("CPU Temperature: "), textStyle(fmt.Sprintf("%.1f°C", m.CPUTemp)))
    s += fmt.Sprintf("%s%s\n", labelStyle("Memory Total: "), textStyle(fmt.Sprintf("%d bits (%.2f GB)", m.MemTotal, float64(m.MemTotal)/math.Pow10(9))))
    s += fmt.Sprintf("%s%s\n", labelStyle("Memory Usage: "), textStyle(fmt.Sprintf("%.1f%%", m.MemPercent)))
    s += fmt.Sprintf("%s%s\n", labelStyle("Disk Total: "), textStyle(fmt.Sprintf("%.2f GB", float64(m.DiskTotal)/math.Pow10(9))))
    diskUsagePercent := 0.0
    if m.DiskTotal > 0 {
        diskUsagePercent = (float64(m.DiskUsed) / float64(m.DiskTotal)) * 100
    }
    s += fmt.Sprintf("%s%s\n", labelStyle("Disk Usage: "), textStyle(fmt.Sprintf("%.1f%%", diskUsagePercent)))
    s += m.progressBar.View() + "\n" // Use the model’s instance of the progress bar
    s += fmt.Sprintf("%s%s\n", labelStyle("Disk Free: "), textStyle(fmt.Sprintf("%.2f GB", float64(m.DiskTotal-m.DiskUsed)/math.Pow10(9))))
    // Uncomment the following lines if GPU stats become available
    // s += fmt.Sprintf("%s%s\n", labelStyle("GPU: "), textStyle(m.GPUName))
    // s += fmt.Sprintf("%s%s\n", labelStyle("GPU Usage: "), textStyle(fmt.Sprintf("%.1f%%", m.GPUPercent)))
    // s += fmt.Sprintf("%s%s\n", labelStyle("GPU Temperature: "), textStyle(fmt.Sprintf("%.1f°C", m.GPUTemp)))
    s += fmt.Sprintf("\n%s\n", helpStyle("Press `q` to quit"))
    s = terminalStyle.Render(s)
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
