package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
)

// CreateDue records a shop due, rent, or bill
func CreateDue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, `{"error": "Unauthorized access"}`, http.StatusUnauthorized)
		return
	}

	var due models.Due
	if err := json.NewDecoder(r.Body).Decode(&due); err != nil {
		http.Error(w, `{"error": "Invalid JSON input"}`, http.StatusBadRequest)
		return
	}

	query := `INSERT INTO dues (user_id, title, amount, due_date) 
	          VALUES ($1, $2, $3, $4) RETURNING id, status, created_at`
	err := database.DB.QueryRow(query, userID, due.Title, due.Amount, due.DueDate).Scan(&due.ID, &due.Status, &due.CreatedAt)
	if err != nil {
		http.Error(w, `{"error": "Failed to record due"}`, http.StatusInternalServerError)
		return
	}

	due.UserID = userID
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Due recorded successfully",
		"due":     due,
	})
}

// GetDues retrieves all shop dues
func GetDues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, `{"error": "Unauthorized access"}`, http.StatusUnauthorized)
		return
	}

	query := `SELECT id, user_id, title, amount, due_date, status, created_at 
	          FROM dues WHERE user_id = $1 ORDER BY due_date ASC`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch dues"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	dues := []models.Due{}
	for rows.Next() {
		var d models.Due
		if err := rows.Scan(&d.ID, &d.UserID, &d.Title, &d.Amount, &d.DueDate, &d.Status, &d.CreatedAt); err != nil {
			http.Error(w, `{"error": "Error scanning due data"}`, http.StatusInternalServerError)
			return
		}
		dues = append(dues, d)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, `{"error": "Error during due iteration"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dues)
}