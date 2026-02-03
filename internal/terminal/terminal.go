package terminal

import (
	"math/rand"
	"time"
)

// Status représente le statut d'un terminal
type Status string

const (
	StatusOK      Status = "OK"
	StatusWarning Status = "Warning"
	StatusError   Status = "Error"
	StatusSyncing Status = "Unsyncing"
)

// Terminal représente un terminal de paiement électronique
type Terminal struct {
	ID       string
	Name     string
	Status   Status
	Battery  int
	LastSync time.Time
}

// GetMockTerminals retourne la liste des terminaux pour la démo
func GetMockTerminals() []Terminal {
	now := time.Now()
	return []Terminal{
		{
			ID:       "TPE-001",
			Name:     "Terminal Narbonne 1",
			Status:   StatusOK,
			Battery:  85,
			LastSync: now.Add(-27 * time.Minute),
		},
		{
			ID:       "TPE-002",
			Name:     "Terminal Toulouse 4",
			Status:   StatusWarning,
			Battery:  42,
			LastSync: now.Add(-3 * time.Minute),
		},
		{
			ID:       "TPE-003",
			Name:     "Terminal Perpignan 7",
			Status:   StatusOK,
			Battery:  27,
			LastSync: now.Add(-11 * time.Minute),
		},
		{
			ID:       "TPE-004",
			Name:     "Terminal Saint Cyprien 2",
			Status:   StatusError,
			Battery:  91,
			LastSync: now.Add(-9 * time.Minute),
		},
		{
			ID:       "TPE-005",
			Name:     "Terminal Rivesalte 1",
			Status:   StatusSyncing,
			Battery:  63,
			LastSync: now.Add(-39 * time.Minute),
		},
		{
			ID:       "TPE-006",
			Name:     "Terminal Thuir 3",
			Status:   StatusOK,
			Battery:  8,
			LastSync: now.Add(-15 * time.Minute),
		},
	}
}

// RandomizeState modifie aléatoirement le statut et la batterie du terminal (pour simulation)
func (t *Terminal) RandomizeStatus() {
	statuses := []Status{StatusOK, StatusWarning, StatusError, StatusSyncing}
	t.Status = statuses[rand.Intn(len(statuses))]

	// Ajuste la batterie valeur entre -10 et +10 et incrémente ou décrémente la batterie
	change := rand.Intn(21) - 10
	t.Battery += change	
	if t.Battery > 100 {
		t.Battery = 100
	}
	if t.Battery < 0 {
		t.Battery = 0
	}

	t.LastSync = time.Now()
}


func (t *Terminal) StatusColor() string {
	switch t.Status {
	case StatusOK :
		return "#00FF00";
	case StatusWarning :
		return "#FFA500";
	case StatusError :
		return "#FF0000";
	case StatusSyncing :
		return "#00BFFF";
	default:
		return "#FFFFFF";
	}
}

func (t *Terminal) StatusEmoji() string {
	switch t.Status {
	case StatusOK :
		return "✅"
	case StatusWarning :
		return "⚠️"
	case StatusError :
		return "❌"
	case StatusSyncing :
		return "🔄"
	default:
		return "❓"
	}
}