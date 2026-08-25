package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"log"

	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
)

func CreateDue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)

	var d models.Due
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	d.Status = "Unpaid"
	query := `INSERT INTO dues (user_id, customer_name, phone, total_due, paid_amount, status) VALUES ($1, $2, $3, $4, 0.00, $5) RETURNING id`
	err := database.DB.QueryRow(query, userID, d.CustomerName, d.Phone, d.TotalDue, d.Status).Scan(&d.ID)
	if err != nil {
		log.Println("=== ADD DUE DB ERROR ===", err)
		http.Error(w, `{"error": "Failed to record due"}`, http.StatusInternalServerError)
		return
	}

	d.UserID = userID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}

func GetDues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)

	rows, err := database.DB.Query(`SELECT id, user_id, customer_name, COALESCE(phone, ''), total_due, paid_amount, status, created_at FROM dues WHERE user_id = $1 ORDER BY id DESC`, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch dues"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	dues := []models.Due{}
	for rows.Next() {
		var d models.Due
		rows.Scan(&d.ID, &d.UserID, &d.CustomerName, &d.Phone, &d.TotalDue, &d.PaidAmount, &d.Status, &d.CreatedAt)
		dues = append(dues, d)
	}

	json.NewEncoder(w).Encode(dues)
}

func PayDueAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)
	dueID := strings.TrimPrefix(r.URL.Path, "/dues/")

	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	var totalDue, currentPaid float64
	err := database.DB.QueryRow(`SELECT total_due, paid_amount FROM dues WHERE id = $1 AND user_id = $2`, dueID, userID).Scan(&totalDue, &currentPaid)
	if err != nil {
		http.Error(w, `{"error": "Due record not found"}`, http.StatusNotFound)
		return
	}

	newPaid := currentPaid + payload.Amount
	newStatus := "Partial"
	if newPaid >= totalDue {
		newPaid = totalDue
		newStatus = "Paid"
	}

	_, err = database.DB.Exec(`UPDATE dues SET paid_amount = $1, status = $2 WHERE id = $3 AND user_id = $4`, newPaid, newStatus, dueID, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to update payment"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Payment recorded successfully"})
}