package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/spinner"
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
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
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

	filledStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")) // Green
	emptyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")) // Grey

	bar := filledStyle.Render(strings.Repeat("█", filled)) +
		emptyStyle.Render(strings.Repeat("░", empty))

	return bar
}
