package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)


/*
	TODO: Change color and styles
	lipgloss.AdaptiveColor{Light: "236", Dark: "248"} //adaptive


	CSS style Styles

	var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    PaddingTop(2).
    PaddingLeft(4).
    Width(22)
	fmt.Println(style.Render("Hello, kitty"))

*/

// Available spinners
var spinners = []spinner.Spinner{
	spinner.Line,
	spinner.Dot,
	spinner.Pulse,
}

// Styles
var (
	barMax = "#fd1259"
	barMin ="#80a5f6"
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#c209f9")).Render
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#80a5f6")).Render
	terminalStyle = lipgloss.NewStyle().Background(lipgloss.Color("#020a1d")).Width(100) //TODO: change colors
	textStyle    = lipgloss.NewStyle().Background(lipgloss.Color("#020a1d")).Foreground(lipgloss.Color("#f2f6ee")).Render
	spinnerStyle = lipgloss.NewStyle().Background(lipgloss.Color("#020a1d")).Foreground(lipgloss.Color("#ff06b5"))
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#917faa")).Render

	// tableStyle := lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("228")).BorderTop(true).BorderLeft(true)
	// tableStyle.Header = table.Header.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240")).BorderBottom(true).Bold(true)
	TableStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#c209f9")).
			BorderTop(true).
			BorderLeft(true).
			BorderBottom(true).
			BorderRight(true)

	TableHeaderStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("#80a5f6")).
				BorderBottom(true).
				Bold(true)
)

// SpinnerModel holds the spinner state
type SpinnerModel struct {
	Index   int
	Spinner spinner.Model
}

// Init initializes the spinner
func (m SpinnerModel) Init() tea.Cmd {
	return m.Spinner.Tick
}

// Update handles spinner state changes
func (m SpinnerModel) Update(msg tea.Msg) (SpinnerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}


// Define the gradient colors for progress bars
var ProgressBar = progress.New(progress.WithScaledGradient(barMin, barMax))


// Function to update the progress bar percentage
func UpdateProgressBar(value float64) tea.Cmd {
	if value < 0 {
		value = 0
	} else if value > 1 {
		value = 1
	}
	return ProgressBar.SetPercent(value)
}

// Function to render the progress bar
func RenderProgressBar() string {
	return ProgressBar.View()
}