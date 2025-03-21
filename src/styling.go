package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/spinner"
	// "github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
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

	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#c209f9")).Render
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#80a5f6")).Render
	terminalStyle = lipgloss.NewStyle().Background(lipgloss.Color("#020a1d")).Width(100) //TODO: change colors
	textStyle    = lipgloss.NewStyle().Background(lipgloss.Color("#020a1d")).Foreground(lipgloss.Color("#f2f6ee")).Render
	spinnerStyle = lipgloss.NewStyle().Background(lipgloss.Color("#020a1d")).Foreground(lipgloss.Color("#0B666A"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
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

// RenderProgressBar generates a progress bar
func RenderProgressBar(percent float64, width int) string {
	if percent < 0 {
		percent = 0
	} else if percent > 100 {
		percent = 100
	}

	filled := int((percent / 100.0) * float64(width))
	if filled > width {
		filled = width
	}
	empty := width - filled
	if empty < 0 {
		empty = 0
	}

	// Gradient colors for progress
	var color string
	switch {
	case percent < 34:
		color = "#f135cc" // Light Blue (Low Usage)
	case percent < 67:
		color = "#c209f9" // Purple (Medium Usage)
	default:
		color = "#fd1259" // Pink/Red (High Usage)
	}
	filledStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
	emptyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#bcd2ab")) // Light Grey

	// Construct the progress bar
	bar := filledStyle.Render(strings.Repeat("█", filled)) +
		emptyStyle.Render(strings.Repeat("░", empty))

	return bar
}

// // Define the gradient colors for progress bars
// var ProgressBar = progress.New(progress.WithScaledGradient("#80a5f6", "#fd1259"))


// // Function to update the progress bar percentage
// func UpdateProgressBar(value float64) tea.Cmd {
// 	if value < 0 {
// 		value = 0
// 	} else if value > 1 {
// 		value = 1
// 	}
// 	return ProgressBar.SetPercent(value)
// }

// // Function to render the progress bar
// func RenderProgressBar() string {
// 	return ProgressBar.View()
// }