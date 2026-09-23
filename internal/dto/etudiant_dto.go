package dto

type CreateEtudiantRequest struct {
	Matricule       string `json:"matricule" validate:"required,max=50"`
	Nom             string `json:"nom" validate:"required,min=2,max=100"`
	Prenom          string `json:"prenom" validate:"required,min=2,max=100"`
	Telephone       string `json:"telephone" validate:"omitempty,max=30"`
	IDEtablissement int64  `json:"id_etablissement" validate:"required"`
	IDFiliere       int64  `json:"id_filiere" validate:"required"`
}

type UpdateEtudiantRequest struct {
	Matricule       string `json:"matricule" validate:"required,max=50"`
	Nom             string `json:"nom" validate:"required,min=2,max=100"`
	Prenom          string `json:"prenom" validate:"required,min=2,max=100"`
	Telephone       string `json:"telephone" validate:"omitempty,max=30"`
	IDEtablissement int64  `json:"id_etablissement" validate:"required"`
	IDFiliere       int64  `json:"id_fiiliere" validaate:"required"`
}

type EtudiantResponse struct {
	ID              int64  `json:"id"`
	Matricule       int64  `json:"matricule"`
	Nom             string `json:"nom"`
	Prenom          string `json:"prenom"`
	Telephone       string `json:"telephone"`
	IDUser          int64  `json:"id_user"`
	IDEtablissement int64  `json:"id_etablissement"`
	IDFiliere       int64  `json:"id_filiere"`
}

type EtudiantDetailResponse struct {
	ID              int64  `json:"id"`
	Matricule       int64  `json:"matricule"`
	Nom             string `json:"nom"`
	Prenom          string `json:"prenom"`
	Telephone       string `json:"telephone"`
	IDUser          int64  `json:"id_user"`
	IDEtablissement int64  `json:"id_etablissement"`
	IDFiliere       int64  `json:"id_filiere"`

	Etablissement *EtablissementResponse `json:"etablissement,omitempty"`
	Filiere       *FiliereResponse       `json:"filiere,omitempty"`
}
