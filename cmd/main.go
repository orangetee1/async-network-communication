package main

import (
	"async_communication/internal/ui"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(ui.InitMain(f), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
