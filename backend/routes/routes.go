package routes

import (
	"net/http"

	"github.com/mdshafiulalamsagar/mystore/backend/controllers"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
)

// SetupRoutes registers all public and protected REST API routes
func SetupRoutes() {
	// Public routes
	http.HandleFunc("/api/signup", controllers.RegisterUser)
	http.HandleFunc("/api/login", controllers.LoginUser)

	// Protected Inventory routes
	http.HandleFunc("/api/inventory/add", middleware.AuthMiddleware(controllers.CreateInventoryItem))
	http.HandleFunc("/api/inventory/list", middleware.AuthMiddleware(controllers.GetInventoryItems))

	// Protected Transaction routes
	http.HandleFunc("/api/transactions/add", middleware.AuthMiddleware(controllers.CreateTransaction))
	http.HandleFunc("/api/transactions/list", middleware.AuthMiddleware(controllers.GetTransactions))
	http.HandleFunc("/api/transactions/summary", middleware.AuthMiddleware(controllers.GetFinancialSummary))

	// Protected Task routes
	http.HandleFunc("/api/tasks/add", middleware.AuthMiddleware(controllers.CreateTask))
	http.HandleFunc("/api/tasks/list", middleware.AuthMiddleware(controllers.GetTasks))

	// Protected Due routes
	http.HandleFunc("/api/dues/add", middleware.AuthMiddleware(controllers.CreateDue))
	http.HandleFunc("/api/dues/list", middleware.AuthMiddleware(controllers.GetDues))
}