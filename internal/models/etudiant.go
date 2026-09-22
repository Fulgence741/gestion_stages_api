package models

type Etudiant struct {
	ID              int64
	Matricule       string
	Nom             string
	Prenom          string
	Telephone       string
	IDUser          int64
	IDEtablissement int64
	IDFiliere       int64
}
