package ui

import "github.com/charmbracelet/lipgloss"

var (
	ErrorStyle      = lipgloss.NewStyle().Background(lipgloss.Color("9")).MarginTop(2)
	WeatherHeader   = lipgloss.NewStyle().Background(lipgloss.Color("5")).SetString("Weather")
	PlaceInfoHeader = lipgloss.NewStyle().Background(lipgloss.Color("5")).SetString("Place information")
)
