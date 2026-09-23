package dto

import "time"

type CreateStageRequest struct {
	IntituleStage string    `json:"intitule_stage" validate:"required,max=200"`
	Description   string    `json:"description"`
	DateDebut     time.Time `json:"date_debut" validate:"required"`
	DateFin       time.Time `json:"date_fin"`
	NiveauEtude   string    `json:"niveau_etude" validate:"required,max=50"`
	IDMaitreStage int64     `json:"id_maitre_stage" validate:"required"`
}

type UpdateStageRequest struct {
	IntituleStage string    `json:"intitule_stage" validate:"required,max=200"`
	Description   string    `json:"description"`
	DateDebut     time.Time `json:"date_debut" validate:"required"`
	DateFin       time.Time `json:"date_fin"`
	NiveauEtude   string    `json:"niveau_etude" validate:"required,max=50"`
	IDMaitreStage int64     `json:"id_maitre_stage" validate:"required"`
}

type SatgeResponse struct {
	ID            int64     `json:"id"`
	IntituleStage string    `json:"intitule_stage"`
	Description   string    `json:"description"`
	DateDebut     time.Time `json:"date_debut"`
	DateFin       time.Time `json:"date_fin"`
	NiveauEtude   string    `json:"niveau_etude"`
	IDEtudiant    int64     `json:"id_etudiant"`
	IDMaitreStage int64     `json:"id_maitre_stage"`
}

type StageDetailResponse struct {
	ID            int64     `json:"id"`
	IntituleStage string    `json:"intitule_stage"`
	Description   string    `json:"description"`
	DateDebut     time.Time `json:"date_debut"`
	DateFin       time.Time `json:"date_fin"`
	NiveauEtude   string    `json:"niveau_etude"`
	IDEtudiant    int64     `json:"id_etudiant"`
	IDMaitreStage int64     `json:"id_maitre_stage"`
}
