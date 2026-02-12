package model

import (
	"math/rand"
	"time"

	"github.com/Mocky-FS/tpe-monitor/internal/terminal"
	"github.com/Mocky-FS/tpe-monitor/internal/view"
	tea "github.com/charmbracelet/bubbletea"
)

// tickMsg est envoyé périodiquement pour le refresh auto
type tickMsg time.Time

// Model représente l'état de l'application
type Model struct {
	terminals []terminal.Terminal
	cursor    int
	quitting  bool
}

func New() Model {
	return Model{
		terminals: terminal.GetMockTerminals(),
		cursor:    0,
	}
}

// Initialise le programme
func (m Model) Init() tea.Cmd {
	return tickCmd()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.terminals)-1 {
				m.cursor++
			}

		case "r":
			// Refresh manuel
			return m, randomizeTerminals(&m)
		}

	case tickMsg:
		// Refresh auto toutes les 10 sec
		return m, tea.Batch(
			randomizeTerminals(&m),
			tickCmd(),
		)
	}

	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return "Au revoir !\n"
	}

	s := view.RenderTitle()

	// Afficher chaque terminal
	for i, t := range m.terminals {
		s += view.RenderTerminal(t, m.cursor == i) + "\n"
	}

	s += view.RenderStatusBar(m.terminals)
	s += view.RenderHelp()

	return s
}

// tickCmd retourne une commande pour le prochain tick
func tickCmd() tea.Cmd {
	return tea.Tick(10*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// randomizeTerminals randomize 1-2 terminaux
func randomizeTerminals(m *Model) tea.Cmd {
	return func() tea.Msg {
		// choisir 1 ou 2 terminaux au hasard
		count := rand.Intn(2) + 1
		for i := 0; i < count; i++ {
			idx := rand.Intn(len(m.terminals))
			m.terminals[idx].RandomizeStatus()
		}
		return nil
	}
}
