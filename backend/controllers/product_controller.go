package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)

	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	query := `INSERT INTO products (user_id, name, price, stock) VALUES ($1, $2, $3, $4) RETURNING id`
	err := database.DB.QueryRow(query, userID, p.Name, p.Price, p.Stock).Scan(&p.ID)
	if err != nil {
		log.Println("=== ADD PRODUCT DB ERROR ===", err)
		http.Error(w, `{"error": "Failed to add product"}`, http.StatusInternalServerError)
		return
	}

	p.UserID = userID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)

	rows, err := database.DB.Query(`SELECT id, user_id, name, price, stock FROM products WHERE user_id = $1 ORDER BY id DESC`, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch products"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.UserID, &p.Name, &p.Price, &p.Stock)
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)
	productID := strings.TrimPrefix(r.URL.Path, "/products/")

	_, err := database.DB.Exec(`DELETE FROM products WHERE id = $1 AND user_id = $2`, productID, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to delete product"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted"})
}