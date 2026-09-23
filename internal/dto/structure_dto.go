package dto

type CreateStructureRequest struct {
	NomStructure string `json:"nom_structure" validate:"required,max=150"`
	Ville        string `json:"ville" validate:"required,max=100"`
	Adresse      string `json:"adresse" validate:"omitempty,max=255"`
	Telephone    string `json:"telephone" validate:"omitempty,max=30"`
	Email        string `json:"email" validate:"omitempty,email,max=255"`
}

type UpdateStructureRequest struct {
	NomStructure string `json:"nom_structure" validate:"required,max=150"`
	Ville        string `json:"ville" validate:"required,max=100"`
	Adresse      string `json:"adresse" validate:"omitempty,max=255"`
	Telephone    string `json:"telephone" validate:"omitempty,max=30"`
	Email        string `json:"email" validate:"omitempty,email,max=255"`
}

type StructureResponse struct {
	ID           int64  `json:"id"`
	NomStructure string `json:"nom_structure"`
	Ville        string `json:"ville"`
	Adresse      string `json:"adresse"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email,omitempty"`
}
