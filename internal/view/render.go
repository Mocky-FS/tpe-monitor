package view

import (
	"fmt"

	"github.com/Mocky-FS/tpe-monitor/internal/terminal"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
)

// Styles globaux
var (
	// Style de la ligne selectionnée
	selectedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1E3A5F")).
			Bold(true)

	// style du titre
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#0066CC")).
			Padding(0, 2)

	// style de la bordure principale
	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#0066CC")).
			Padding(0, 1)

	// style de la barre de statut
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))

	// styles des statuts
	okStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Bold(true)
	warningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Bold(true)
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4444")).Bold(true)
	syncingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Bold(true)
)

// RenderTitle affiche le titre de l'application
func RenderTitle() string {
	return titleStyle.Render("🖥️  TPE Monitor") + "\n\n"
}

// RenderTerminal affiche une ligne terminal
func RenderTerminal(t terminal.Terminal, selected bool) string {
	// formater le statut avec sa couleur
	status := renderStatus(t)

	// formater la batterie
	battery := renderBattery(t.Battery)

	// formater le temps
	lastSync := humanize.Time(t.LastSync)

	// construire la ligne
	line := fmt.Sprintf("%-8s  %-25s  %-18s  %-6s  %s",
		t.ID,
		t.Name,
		status,
		battery,
		lastSync,
	)

	// appliquer le style si selectionné
	if selected {
		return "▶ " + selectedStyle.Render(line)
	}
	return " " + line
}

// RenderStatusBar affiche la barre de stats en bas
func RenderStatusBar(terminals []terminal.Terminal) string {
	ok, warning, err, syncing := countByStatus(terminals)

	bar := fmt.Sprintf(
		" %d terminaux  •  %s  •  %s  •  %s  •  %s",
		len(terminals),
		okStyle.Render(fmt.Sprintf("%d OK", ok)),
		warningStyle.Render(fmt.Sprintf("%d Warning", warning)),
		errorStyle.Render(fmt.Sprintf("%d Error", err)),
		syncingStyle.Render(fmt.Sprintf("%d Syncing", syncing)),
	)
	return "\n" + statusBarStyle.Render(bar)
}

// RenderHelp affiche les touches disponibles
func RenderHelp() string {
	return statusBarStyle.Render("\n  ↑↓ naviguer  •  r refresh  •  q quitter")
}

// renderStatus retourne le statut coloré
func renderStatus(t terminal.Terminal) string {
	label := fmt.Sprintf("%s %-8s", t.StatusEmoji(), t.Status)
	switch t.Status {
	case terminal.StatusOK:
		return okStyle.Render(label)
	case terminal.StatusWarning:
		return warningStyle.Render(label)
	case terminal.StatusError:
		return errorStyle.Render(label)
	case terminal.StatusSyncing:
		return syncingStyle.Render(label)
	default:
		return label
	}
}

// renderBattery retourne la batterie avec la couleur selon le niveau
func renderBattery(level int) string {
	icon := "🔋"
	label := fmt.Sprintf("%s %d%%", icon, level)

	switch {
	case level > 50:
		return okStyle.Render(label)
	case level > 20:
		return warningStyle.Render(label)
	default:
		return errorStyle.Render(label)
	}
}

// countByStatus compte les terminaux par statut
func countByStatus(terminals []terminal.Terminal) (ok, warning, err, syncing int) {
	for _, t := range terminals {
		switch t.Status {
		case terminal.StatusOK:
			ok++
		case terminal.StatusWarning:
			warning++
		case terminal.StatusError:
			err++
		case terminal.StatusSyncing:
			syncing++
		}
	}
	return
}
