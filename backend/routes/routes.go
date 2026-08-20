package routes

import (
	"net/http"

	"github.com/mdshafiulalamsagar/mystore/backend/controllers"
)

// SetupRoutes registers API routes
func SetupRoutes() {
	http.HandleFunc("/api/signup", controllers.RegisterUser)
	http.HandleFunc("/api/login", controllers.LoginUser)
}