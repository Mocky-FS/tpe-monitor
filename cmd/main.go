package main

import (
	"fmt"
	"os"

	"github.com/Mocky-FS/tpe-monitor/internal/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := model.New()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Erreur: %v\n", err)
		os.Exit(1)
	}
}
