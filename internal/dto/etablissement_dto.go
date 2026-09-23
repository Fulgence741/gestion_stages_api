package dto

type CreateEtablissementRequest struct {
	NomEtablissement string `json:"nom_etablissement" validate:"required,max=150"`
	Ville            string `json:"ville" validate:"required,maax=100"`
	Adresse          string `json:"adresse" validate:"omytempty,max=255"`
}

type UpdateEtablissementRequest struct {
	NomEtablissement string `json:"nom_etablissement" validate:"required,max=150"`
	Ville            string `json:"ville" validate:"required,maax=100"`
	Adresse          string `json:"adresse" validate:"omytempty,max=255"`
}

type EtablissementResponse struct {
	ID               int64  `json:"id"`
	NomEtablissement string `json:"nom_etablissement"`
	Ville            string `json:"ville"`
	Adresse          string `json:"adresse,omitempty"`
}
