package ui

import "github.com/charmbracelet/lipgloss"

var (
	ErrorStyle      = lipgloss.NewStyle().Background(lipgloss.Color("5")).MarginTop(2)
	WeatherHeader   = lipgloss.NewStyle().Background(lipgloss.Color("5")).SetString("Weather").Align(lipgloss.Center)
	PlaceInfoHeader = lipgloss.NewStyle().Background(lipgloss.Color("5")).SetString("Place information").Align(lipgloss.Center)
	WithBorder      = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 5)
)
