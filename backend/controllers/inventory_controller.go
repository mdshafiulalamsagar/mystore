package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
)

// CreateInventoryItem adds a new stock item for the authenticated user
func CreateInventoryItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, `{"error": "Unauthorized access"}`, http.StatusUnauthorized)
		return
	}

	var item models.InventoryItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, `{"error": "Invalid JSON input"}`, http.StatusBadRequest)
		return
	}

	query := `INSERT INTO inventory (user_id, item_name, quantity, unit_price, low_stock_limit) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	err = database.DB.QueryRow(query, userID, item.ItemName, item.Quantity, item.UnitPrice, item.LowStockLimit).Scan(&item.ID, &item.CreatedAt)
	if err != nil {
		http.Error(w, `{"error": "Failed to create inventory item"}`, http.StatusInternalServerError)
		return
	}

	item.UserID = userID
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Inventory item added successfully",
		"item":    item,
	})
}

// GetInventoryItems fetches all stock items for the authenticated user
func GetInventoryItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, `{"error": "Unauthorized access"}`, http.StatusUnauthorized)
		return
	}

	query := `SELECT id, user_id, item_name, quantity, unit_price, low_stock_limit, created_at 
	          FROM inventory WHERE user_id = $1 ORDER BY id DESC`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch inventory items"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	items := []models.InventoryItem{}
	for rows.Next() {
		var item models.InventoryItem
		err := rows.Scan(&item.ID, &item.UserID, &item.ItemName, &item.Quantity, &item.UnitPrice, &item.LowStockLimit, &item.CreatedAt)
		if err != nil {
			http.Error(w, `{"error": "Error scanning inventory data"}`, http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, `{"error": "Error during inventory row iteration"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}