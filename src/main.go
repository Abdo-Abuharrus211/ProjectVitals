package main

import (
	"fmt"
	"math"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
    progress "github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
)

const refreshInterval = time.Second // Update interval

// Model struct to store app state
type model struct {
	cursor     int
	spinner    SpinnerModel
    diskBar progress.Model
    memoryBar progress.Model
	statsTable table.Model

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

// Initialize the model with Bubble Tea components
func initialModel() model {
    dbar := progress.New(progress.WithScaledGradient(barMin, barMax))
    mbar := progress.New(progress.WithScaledGradient(barMin, barMax))
	spin := spinner.New()
	spin.Spinner = spinner.Dot // Default spinner type
	spin.Style = spinnerStyle
	columns := []table.Column{
        {Title: "Stat", Width: 20},
        {Title: "Value", Width: 30},
    }
    t := table.New(
        table.WithColumns(columns),
        table.WithFocused(true),
        table.WithHeight(10),
    )
    customStyles := table.Styles{
        Header:   TableHeaderStyle,
        Selected: lipgloss.NewStyle().Foreground(lipgloss.Color("#f2f6ee")),
        Cell:     lipgloss.NewStyle().Foreground(lipgloss.Color("#f2f6ee")),
	}
	t.SetStyles(customStyles)

	return model{
        cursor:      0,
        spinner:     SpinnerModel{Spinner: spin},
        diskBar: dbar,
        memoryBar: mbar,
		statsTable: t,
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

/*
	Update the Stats Table on each refresh, simply rewrite the model's latest data
*/
func (m *model) updateStatsTable() {
    rows := []table.Row{
        {"PC Name", m.PCName},
        {"OS", m.OSName + " " + m.OSVersion},
        {"CPU", m.CPUName},
        {"CPU Usage", fmt.Sprintf("%.1f%%", m.CPUPercent)},
        {"CPU Temperature", fmt.Sprintf("%.1f°C", m.CPUTemp)},
        {"Memory Total", fmt.Sprintf("%.2f GB", float64(m.MemTotal)/math.Pow10(9))},
        {"Disk Total", fmt.Sprintf("%.2f GB", float64(m.DiskTotal)/math.Pow10(9))},
        {"Disk Free", fmt.Sprintf("%.2f GB", float64(m.DiskTotal-m.DiskUsed)/math.Pow10(9))},
    }
    m.statsTable.SetRows(rows)
}


/*
	 Update function handles messages and interactions, in turn updates TUI
*/
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
        diskProgressCmd := m.diskBar.SetPercent(diskUsage)
        memoryUsage := m.MemPercent / 100.0
        memoryProgressCmd := m.memoryBar.SetPercent(memoryUsage)
		m.updateStatsTable()
        return m, tea.Batch(refresh(), diskProgressCmd, memoryProgressCmd)

    case progress.FrameMsg:
        diskProgressModel, diskProgressCmd := m.diskBar.Update(msg)
        m.diskBar = diskProgressModel.(progress.Model)
        memProgressModel, memProgressCmd := m.memoryBar.Update(msg)
        m.memoryBar = memProgressModel.(progress.Model)
        return m, tea.Batch(diskProgressCmd, memProgressCmd)
    
        
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
	s := fmt.Sprintf("%s\n\n", titleStyle("ProjectVitalis"))
    s += fmt.Sprintf("%s%s\n\n", labelStyle("Monitoring system vitals "), m.spinner.Spinner.View())
    s += TableStyle.Render(m.statsTable.View()) + "\n\n"

    s += fmt.Sprintf("%s%s\n", labelStyle("Memory Usage: "), textStyle(fmt.Sprintf("%.1f%%", m.MemPercent)))
    s += m.memoryBar.View() + "\n\n"
    
    diskUsagePercent := 0.0
    if m.DiskTotal > 0 {
        diskUsagePercent = (float64(m.DiskUsed) / float64(m.DiskTotal)) * 100
    }
    s += fmt.Sprintf("%s%s\n", labelStyle("Disk Usage: "), textStyle(fmt.Sprintf("%.1f%%", diskUsagePercent)))
    s += m.diskBar.View() + "\n"
    
    s += fmt.Sprintf("\n%s\n", helpStyle("Press `q` to quit"))
    return terminalStyle.Render(s)
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
