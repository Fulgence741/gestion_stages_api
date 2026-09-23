package dto

type CreateFiliereRequest struct {
	NomFiliere string `json:"nom_filere" validate:"required,max=255"`
}

type UpdateFilereRequest struct {
	NomFiliere string `json:"nom_filere" validate:"required,max=255"`
}

type FiliereResponse struct {
	ID         int64  `json:"id"`
	NomFiliere string `json:"nom_filiere"`
}
