package dto

import "time"

type UserResponse struct {
	ID           int64     `json:"id"`
	Nom          string    `json:"nom"`
	Prenom       string    `json:"prenom"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	Statut       string    `json:"statut"`
	DateCreation time.Time `json:"date_creation"`
}

type UpdateUserRequest struct {
	Nom    string `json:"nom" validate:"required,min=2,max=100"`
	Prenom string `json:"prenom" validate:"required,min=2,max=100"`
	Email  string `json:"email" validate:"required,min=2,max=255"`
}

type UpdateStatutRequest struct {
	Statut string `json:"statut" validate:"required"`
}

type ChangePasswordRequest struct {
	AncienMotDePasse  string `json:"ancien_mot_de_passe" validate:"required"`
	NouveauMotDePasse string `json:"nouveau_mot_de_passe" validate:"required,min=8,max=255"`
}
