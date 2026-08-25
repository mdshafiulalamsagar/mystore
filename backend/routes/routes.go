package routes

import (
	"net/http"

	"github.com/mdshafiulalamsagar/mystore/backend/controllers"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
)

func SetupRoutes() {
	http.HandleFunc("/signup", controllers.RegisterUser)
	http.HandleFunc("/login", controllers.LoginUser)

	// Protected Product Routes
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			middleware.Authenticate(controllers.CreateProduct)(w, r)
		} else if r.Method == http.MethodGet {
			middleware.Authenticate(controllers.GetProducts)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			middleware.Authenticate(controllers.DeleteProduct)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			middleware.Authenticate(controllers.CreateTransaction)(w, r)
		} else if r.Method == http.MethodGet {
			middleware.Authenticate(controllers.GetTransactions)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/transactions/summary", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			middleware.Authenticate(controllers.GetTransactionSummary)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Protected Order Routes
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			middleware.Authenticate(controllers.CreateOrder)(w, r)
		} else if r.Method == http.MethodGet {
			middleware.Authenticate(controllers.GetOrders)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut || r.Method == http.MethodPatch {
			middleware.Authenticate(controllers.UpdateOrderStatus)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Protected Due Routes
	http.HandleFunc("/dues", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			middleware.Authenticate(controllers.CreateDue)(w, r)
		} else if r.Method == http.MethodGet {
			middleware.Authenticate(controllers.GetDues)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/dues/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut || r.Method == http.MethodPatch {
			middleware.Authenticate(controllers.PayDueAmount)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
	})
}