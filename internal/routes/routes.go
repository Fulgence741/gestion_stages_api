package routes

import (
	"gestion_stages_api/internal/handlers"
	"net/http"
)

func RegisterRoutes() {

	// AUTHENTIFICATION

	http.HandleFunc("POST /api/auth/register", handlers.AuthRegister)
	http.HandleFunc("POST /api/auth/login", handlers.AuthLogin)
	http.HandleFunc("POST /api/auth/me", handlers.AuthMe)
	http.HandleFunc("POST /api/auth/logout", handlers.AuthLogout)

	// UTILISATEURS

	http.HandleFunc("POST /api/users/create", handlers.CreateUsers)
	http.HandleFunc("POST /api/users/list", handlers.ListUsers)
	http.HandleFunc("POST /api/users/get/{id}", handlers.GetUsers)
	http.HandleFunc("POST /api/users/update/{id}", handlers.UpdateUsers)
	http.HandleFunc("POST /api/users/delete/{id}", handlers.DeleteUsers)
	http.HandleFunc("POST /api/users/change-statut/{id}", handlers.StatutChange)
	http.HandleFunc("POST /api/users/change-password/{id}", handlers.PasswordChange)

	// ETUDIANTS

	http.HandleFunc("POST /api/etudiants/create", handlers.CreateEtudiant)
	http.HandleFunc("POST /api/etudiants/list", handlers.ListEtudiant)
	http.HandleFunc("POST /api/etudiants/get/{id}", handlers.Getetudiant)
	http.HandleFunc("POST /api/etudiants/update/{id}", handlers.UpdateEtudiant)
	http.HandleFunc("POST /api/etudiants/delete/{id}", handlers.DeleteEudiant)
	http.HandleFunc("POST /api/etudiants/stage/{id}", handlers.GetEtudiantStatge)

	// ETABLISSEMENTS

	http.HandleFunc("POST /api/etablissement/create", handlers.CreateEtablissement)
	http.HandleFunc("POST /api/etablissement/list", handlers.ListEtablissement)
	http.HandleFunc("POST /api/etablissement/get/{id}", handlers.GetEtablissement)
	http.HandleFunc("POST /api/etablissement/update/{id}", handlers.UpdateEtablissement)
	http.HandleFunc("POST /api/etablissement/delete/{id}", handlers.DeleteEtablissement)

	// Relation Etablissement - Filiere
	http.HandleFunc("POST /api/etablissement/filieres/{id}", handlers.ListerFilieresEtab)
	http.HandleFunc("POST /api/etablissement/filieres/{id}/add/{idfiliere}", handlers.AddFiliereEtab)
	http.HandleFunc("DELETE /api/etablissement/filieres/{id}/delete/{idfiliere}", handlers.DeleteFiliereEtab)

	// FILIERES

	http.HandleFunc("POST /api/filieres/create", handlers.CreateFiliere)
	http.HandleFunc("POST /api/filieres/list", handlers.ListFilieres)
	http.HandleFunc("POST /api/filieres/get/{id}", handlers.GetFiliere)
	http.HandleFunc("POST /api/filieres/update/{id}", handlers.UpdateFiliere)
	http.HandleFunc("POST /api/filieres/delete/{id}", handlers.DeleteFiliere)

	// STRUCTURES

	http.HandleFunc("POST /api/structure/create", handlers.CreateStructure)
	http.HandleFunc("POST /api/structure/list", handlers.ListStructures)
	http.HandleFunc("POST /api/structure/get/{id}", handlers.GetStructure)
	http.HandleFunc("POST /api/structure/update/{id}", handlers.UpdateStructure)
	http.HandleFunc("POST /api/structure/delete/{id}", handlers.DeleteStructure)

	// Maitres de stage d'une structure
	http.HandleFunc("POST /api/structure/maitre-stage/{id}", handlers.ListMaitreStageStructure)

	// MAITRES DE STAGE

	http.HandleFunc("POST /api/maitre-stage/create", handlers.CreateMaitreSatge)
	http.HandleFunc("POST /api/maitre-stage/list", handlers.ListMaitreStage)
	http.HandleFunc("POST /api/maitre-stage/get/{id}", handlers.GetMaitreStage)
	http.HandleFunc("POST /api/maitre-stage/update/{id}", handlers.UpdateMaitrestage)
	http.HandleFunc("POST /api/maitre-stage/delete/{id}", handlers.DeleteMaitreStage)

	// Stages encadrés par un maître de stage
	http.HandleFunc("POST /api/maitre-stage/stage/{id}", handlers.StageMaitreStage)

	// STAGES

	http.HandleFunc("POST /api/stage/create", handlers.CreateStage)
	http.HandleFunc("POST /api/stage/list", handlers.ListStages)
	http.HandleFunc("POST /api/stage/get/{id}", handlers.GetStage)
	http.HandleFunc("POST /api/stage/update/{id}", handlers.UpdateStage)
	http.HandleFunc("POST /api/stage/delete/{id}", handlers.DeleteStage)

	// RAPPORTS

	http.HandleFunc("POST /api/rapport/create", handlers.CreateRapport)
	http.HandleFunc("POST /api/rapport/list", handlers.ListRapports)
	http.HandleFunc("POST /api/rapport/get/{id}", handlers.GetRapport)
	http.HandleFunc("POST /api/rapport/update/{id}", handlers.UpdateRapport)
	http.HandleFunc("POST /api/rapport/delete/{id}", handlers.DeleteRapport)

	// Rapport d'un stage
	http.HandleFunc("POST /api/rapport/stage/{id}", handlers.GetStageRapport)

	// EVALUATIONS

	http.HandleFunc("POST /api/evaluation/create", handlers.CreateEvaluation)
	http.HandleFunc("POST /api/evaluation/list", handlers.ListEvaluations)
	http.HandleFunc("POST /api/evaluation/get/{id}", handlers.GetEvaluation)
	http.HandleFunc("POST /api/evaluation/update/{id}", handlers.UpdateEvaluation)
	http.HandleFunc("POST /api/evaluation/delete/{id}", handlers.DeleteEvaluation)

	// Evaluations d'un rapport
	http.HandleFunc("POST /api/evaluation/rapport/{id}", handlers.ListRapportEvaluations)
}
