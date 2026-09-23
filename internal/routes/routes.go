package routes

import (
	"gestion_stages_api/internal/handlers"
	"net/http"
)

func RegisterRoutes() {

	// Authentication
	http.Handle("POST /api/auth/register", http.HandlerFunc(handlers.AuthLogin))
	http.Handle("POST /api/auth/login", http.HandlerFunc(handlers.AuthRegister))
	http.Handle("POST /api/auth/me", http.HandlerFunc(handlers.AuthMe))
	http.Handle("POST /api/auth/logout", http.HandlerFunc(handlers.AuthLogout))

	// Users
}
