package models

import (
	"time"
)

type Stage struct {
	ID            int64
	IntituleStage string
	Description   string
	DateDebut     time.Time
	DateFin       *time.Time
	NiveauEtude   string
	IDEtudiant    int64
	IDMaitreStage int64
}
