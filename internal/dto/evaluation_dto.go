package dto

import "time"

type CreateEvaluationRequest struct {
	Note        float64 `json:"note" validate:"required,min=0,maxx=20"`
	Commentaire string  `json:"commentaire"`
}

type UpdateEvaluationRequest struct {
	Note        float64 `json:"note" validate:"required,min=0,maxx=20"`
	Commentaire string  `json:"commentaire"`
}

type EvaluationResponse struct {
	ID             int64     `json:"id"`
	Note           float64   `json:"note"`
	Commentaire    string    `json:"commentaire,omitempty"`
	DateEvaluation time.Time `json:"date_evaluation"`
	IDRapport      int64     `json:"id_rapport"`
	IdUser         int64     `json:"id_user"`
}

type EvaluationDetailResponse struct {
	ID             int64         `json:"id"`
	Note           float64       `json:"note"`
	Commentaire    string        `json:"commentaire,omitempty"`
	DateEvaluation time.Time     `json:"date_evaluation"`
	IDRapport      int64         `json:"id_rapport"`
	IdUser         int64         `json:"id_user"`
	Evaluateur     *UserResponse `json:"evaluateur,omitempty"`
}
