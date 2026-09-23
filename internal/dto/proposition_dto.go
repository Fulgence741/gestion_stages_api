package dto

import "time"

type CreatePropositionRequest struct {
	IDEtablissement int64 `json:"id_etablissement" validate:"required"`
	IDFiliere       int64 `json:"id_filiere" validate:"required"`
}

type PropositionResponse struct {
	IDEtablissement int64     `json:"id_etablissement"`
	IDFiliere       int64     `json:"id_filiere"`
	DateProposition time.Time `json:"date_proposition"`
}

type EtablissementFilieresResponse struct {
	Etablissement EtablissementResponse `json:"etablissement"`
	Filieres      []FiliereResponse     `json:"filieres"`
}

type FiliereEtablissements struct {
	Filiere        FiliereResponse         `json:"filiere"`
	Etablissements []EtablissementResponse `json:"etablissement"`
}
