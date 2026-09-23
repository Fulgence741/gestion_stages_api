package auth

type LoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	MotDePasse string `json:"mot_de_passe" validate:"required"`
}

type LoginResponse struct {
	Token string           `json:"token"`
	User  UserAuthResponse `json:"user"`
}

type UserAuthResponse struct {
	ID     int64  `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Statut string `json:"statut"`
}
