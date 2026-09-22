package main

import (
	"log"
	"net/http"

	"gestion_stages_api/internal/config"
	"gestion_stages_api/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.ConnexionDB(cfg)
	if err != nil {
		log.Fatal("Erreur de la connexion à la base de données :", err)
	}

	defer db.Close()

	log.Println("Connexion à PostgreSQL réussie !!")
	log.Println("Serveur démarré sur le port", cfg.ServerPort)

	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		log.Fatal("Erreur du serveur :", err)
	}
}
