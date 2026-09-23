package dto

import "time"

type CreateRapportRequest struct {
	TitreRapport  string `json:"titre_rapport" validate:"required,max=200"`
	CheminFichier string `json:"chemin_fichier" validate:"required,max=500"`
}

type UpdateRapportRequest struct {
	TitreRapport  string `json:"titre_rapport" validate:"required,max=200"`
	CheminFichier string `json:"chemin_fichier" validate:"required,max=500"`
}

type RapportResponse struct {
	ID            int64     `json:"id"`
	TitreRapport  string    `json:"titre_rapport"`
	CheminFichier string    `json:"chemin_fichier"`
	DateDepot     time.Time `json:"date_depot,omitempty"`
	Statut        string    `json:"statut"`
	NoteFinale    *float64  `json:"note_finale"`
	IDStage       int64     `json:"id_stage"`
}

type RapportDetailResponse struct {
	ID            int64                `json:"id"`
	TitreRapport  string               `json:"titre_rapport"`
	CheminFichier string               `json:"chemin_fichier"`
	DateDepot     time.Time            `json:"date_depot,omitempty"`
	Statut        string               `json:"statut"`
	NoteFinale    *float64             `json:"note_finale"`
	IDStage       int64                `json:"id_stage"`
	Stage         *StageResponse       `json:"stage,omitempty"`
	Evaluation    []EvaluationResponse `json:"evaluation,omitempty"`
}
