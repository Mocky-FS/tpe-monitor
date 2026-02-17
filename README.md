# 🖥️ TPE Monitor

Dashboard TUI pour le monitoring de terminaux de paiement.

![Démo](images/demo.gif)

## 🚀 Installation & Lancement
```bash
# Cloner et lancer
git clone https://github.com/Mocky-FS/tpe-monitor.git
cd tpe-monitor
go run cmd/main.go
```

## 📦 Téléchargement

Binaires précompilés disponibles dans la [Release v1.0.0](https://github.com/Mocky-FS/tpe-monitor/releases/tag/v1.0.0) :

| Plateforme | Fichier |
|------------|---------|
| Linux | `tpe-monitor-linux-amd64` |
| macOS Intel | `tpe-monitor-macos-amd64` |
| macOS Apple Silicon | `tpe-monitor-macos-arm64` |
| Windows | `tpe-monitor-windows-amd64.exe` |

## ⌨️ Utilisation

- `↑↓` : Naviguer entre les terminaux
- `Enter` : Vue détaillée du terminal sélectionné
- `ESC` : Fermer la vue détaillée
- `r`  : Refresh manuel
- `q`  : Quitter

Auto-refresh toutes les 10 secondes.

## 🛠️ Stack

- Go 1.25
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling & colours
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - Spinner component
- [go-humanize](https://github.com/dustin/go-humanize) - Date formatting

## 📁 Structure
```
cmd/main.go              # Point d'entrée
internal/model/          # Logique Bubble Tea
internal/terminal/       # Données des terminaux
internal/view/           # Rendu et styling Lipgloss
```

## 🎯 Contexte

Projet démo pour candidature chez Afsol Perpignan (solutions mPOS/monitoring de flottes TPE).

---

**Auteur** : [Mocky-FS](https://github.com/Mocky-FS)
