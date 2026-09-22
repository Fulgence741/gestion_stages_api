package models

import (
	"time"
)

type User struct {
	ID           int64
	Nom          string
	Prenom       string
	email        string
	MotDePasse   string
	Role         string
	Statut       string
	DateCreation time.Time
}
