package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Standard color tokens for the design system.
	ColorPrimary = lipgloss.Color("63")  // Premium Purple/Blue
	ColorSuccess = lipgloss.Color("42")  // Vibrant Green
	ColorError   = lipgloss.Color("196") // Vibrant Red
	ColorGray    = lipgloss.Color("240")
	ColorWhite   = lipgloss.Color("255")
)
var (
	// HeaderStyle defines the visual style for table headers.
	HeaderStyle = lipgloss.NewStyle().
			Foreground(ColorWhite).
			Background(ColorPrimary).
			Bold(true).
			Padding(0, 1)

	RowStyle = lipgloss.NewStyle().
			Padding(0, 1)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	StatusOkText   = SuccessStyle.Render("ONLINE")
	StatusFailText = ErrorStyle.Render("OFFLINE")
)
