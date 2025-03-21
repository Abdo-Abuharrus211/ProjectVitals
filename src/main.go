package main

import(
	"fmt"
	"os"
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

const refreshInterval = time.Second // this is 1 sec


/*
	Send a refreshMsg once every set interval.
	Returns a Cmd to tick
*/
func refresh() tea.Cmd{
	return tea.Tick(refreshInterval, func(time.Time) tea.Msg{
		return refreshMsg{}
	})
}


/*
	The model struct storesd the app's initial state - Bubbletea boilerplate
*/
type model struct {
	stats []string // items we're concerned about
	cursor	int // what the curso points at
	selected map[int]struct{} // what we've selected from our stats
	PCName string
	OSName string
	OSVersion string
	CPUName string
	CPUPercent float64
	CPUTemp float64
    MemTotal uint64
	MemPercent float64
	DiskTotal uint64
	DiskUsed uint64
	GPUName string
	GPUPercent float64
	GPUTemp float64
}

func initialModel() model{
	return model{
		//our choices of rhings
		stats: []string{}, // empty for now
		selected: make(map[int]struct{}), // the keys refer to the indexes of the `choices` slice above.
	}
}


/*
	Init can return a Cmd that could perform some initial I/O. 
*/
func (m model) Init() tea.Cmd{
	return refresh() // this means no I/O returned at the moment
}


type refreshMsg struct{}

/*
  Updates state when changes occur and updates the app's model in turn.
  Can return a Cmd to invoke more actions.
*/
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
		//TODO: Add the cases for the different options.
    case tea.KeyMsg:
        if msg.String() == "q" {
            return m, tea.Quit // instructs BubbleTea to quit
		}
    case refreshMsg:
		// return model{
		// 	choices:    m.choices,
		// 	cursor:     m.cursor,
		// 	selected:   m.selected,
		// 	CPUPercent: cpu,
		// 	MemPercent: mem,
		// }, tea.Tick(refreshInterval, refresh) // This is live loading or something?
		m.PCName, _ = GetPCName()
		cpu, _ := GetCPUUsage()
		cpuTemp, _ := GetCPUStats()
		memTotal, memUse, _ := GetMemoryStats()
		diskTotal, diskUsed, _ := GetDiskStats()
		
		gpuName,gpuPercent, gpuTemp, _ := GetGPUStats()
		osName, osVersion, _ := GetOSInfo()
		m.CPUName, _ = GetCPUName()

		m.CPUPercent = cpu
		m.CPUTemp = cpuTemp
		m.MemTotal = memTotal // TODO: Convert to MB
		m.MemPercent = memUse
		m.DiskTotal = diskTotal
		m.DiskUsed = diskUsed
		m.GPUName = gpuName
		m.GPUPercent = gpuPercent
		m.GPUTemp = gpuTemp
		m.OSName = osName
		m.OSVersion = osVersion
		return m, refresh() // Schedule next refresh

		
	}

	// Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

/*
	Render UI in terminal.
	Parses the model's state and stringies it, the string becomes the UI.
	BubbleTea handles all the redrawing.
*/
func (m model) View() string {
	// TODO: add the logic for how we render stuff.
    s := "SYSTEM MONITOR\n\n"
    s += fmt.Sprintf("CPU Usage: %.1f%%\n", m.CPUPercent)
    s += fmt.Sprintf("Memory: %d\n", m.MemTotal)
    s += fmt.Sprintf("Memory Usage: %.1f%%\n", m.MemPercent)
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