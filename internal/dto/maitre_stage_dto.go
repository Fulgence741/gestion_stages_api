package dto

type CreateMaireStageRequest struct {
	Nom         string `json:"nom" validate:"required,min=2,max=100"`
	Prenom      string `json:"prenom" validate:"required,min=2,max=100"`
	Fonction    string `json:"fonction" validate:"required,min=2,max=100"`
	Telephone   string `json:"telephone" validate:"omitempty,max=30"`
	Email       string `json:"email" validate:"omitempty,email,max=255"`
	IDStructure int64  `json:"id_structure" validate:"required"`
}

type UpdateMaitreStageRequest struct {
	Nom         string `json:"nom" validate:"required,min=2,max=100"`
	Prenom      string `json:"prenom" validate:"required,min=2,max=100"`
	Fonction    string `json:"fonction" validate:"required,min=2,max=100"`
	Telephone   string `json:"telephone" validate:"omitempty,max=30"`
	Email       string `json:"email" validate:"omitempty,email,max=255"`
	IDStructure int64  `json:"id_structure" validate:"required"`
}

type MaitreStageResponse struct {
	ID          int64  `json:"id"`
	Nom         int64  `json:"nom"`
	Prenom      int64  `json:"prenom"`
	Fonction    string `json:"fonction"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	IDStructure int64  `json:"id_structure"`
}

type MaitreSategDetailRespose struct {
	ID          int64              `json:"id"`
	Nom         int64              `json:"nom"`
	Prenom      int64              `json:"prenom"`
	Fonction    string             `json:"fonction"`
	Telephone   string             `json:"telephone"`
	Email       string             `json:"email"`
	IDStructure int64              `json:"id_structure"`
	Structure   *StructureResponse `json:"structure,omitempty"`
}
