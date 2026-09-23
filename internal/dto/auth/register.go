package auth

type RegisterRequest struct {
	Nom        string `json:"nom" validate:"required,min=2,max=100"`
	Prenom     string `json:"prenom" validate:"required,min=2,max=100"`
	Email      string `json:"email" validate:"required,email,max=255"`
	MotDePasse string `json:"mot_de_passe" validate:"required,min=8,max=255"`
}

type RegisterResponse struct {
	ID     int64  `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Statut string `json:"statut"`
}
