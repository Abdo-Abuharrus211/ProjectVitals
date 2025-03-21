package main

import(
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)


/*
	The model struct storesd the app's initial state - Bubbletea boilerplate
*/
type model struct {
	stats []string // items we're concerned about
	cursor	int // what the curso points at
	selected map[int]struct{} // what we've selected from our stats
	CPUPercent float64
    MemPercent float64
}

func initialModel() model{
	return model{
		//our choices of rhings
		choices: []string{} // empty for now
		selected: make(map[int]struct{}) // the keys refer to the indexes of the `choices` slice above.
	}
}



/*
	Init can return a Cmd that could perform some initial I/O. 
*/
func (m model) Init() tea.cmd{
	return nil // this means no I/O returned at the moment
}

func (m model) View() string {
    s := "SYSTEM MONITOR\n\n"
    s += fmt.Sprintf("CPU Usage: %.1f%%\n", m.CPUPercent)
    s += fmt.Sprintf("Memory Usage: %.1f%%\n", m.MemPercent)
    s += "\nPress q to quit\n"
    return s
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
		return model{
			choices:    m.choices,
			cursor:     m.cursor,
			selected:   m.selected,
			CPUPercent: cpu,
			MemPercent: mem,
		}, tea.Tick(refreshInterval, refresh) // This is live loading or something?
		
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
	// add the logic for how we render stuff.
}

/*
	Main function to run the program.
*/
func main(){
	monitorProgram := tea.NewProgram(initialModel())
	if _, err := monitorProgram.Run(); err != nil {
		fmt.Printf("Error occured: %v", err)
		os.exit(1)
	}
}