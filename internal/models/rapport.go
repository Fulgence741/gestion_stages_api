package models

import (
	"time"
)

type Rapport struct {
	ID            int64
	TitreRapport  string
	CheminFichier string
	DateDepot     *time.Time
	Statut        string
	NoteFinale    *float64
	IDStage       int64
}
