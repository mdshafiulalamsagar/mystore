package routes

import (
	"net/http"

	"github.com/mdshafiulalamsagar/mystore/backend/controllers"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
)

// SetupRoutes registers public and protected API routes
func SetupRoutes() {
	// Public routes
	http.HandleFunc("/api/signup", controllers.RegisterUser)
	http.HandleFunc("/api/login", controllers.LoginUser)

	// Protected Inventory routes wrapped with AuthMiddleware
	http.HandleFunc("/api/inventory/add", middleware.AuthMiddleware(controllers.CreateInventoryItem))
	http.HandleFunc("/api/inventory/list", middleware.AuthMiddleware(controllers.GetInventoryItems))
}