package controllers

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)

	var t models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	query := `INSERT INTO transactions (user_id, type, amount, description) VALUES ($1, $2, $3, $4) RETURNING id`
	err := database.DB.QueryRow(query, userID, t.Type, t.Amount, t.Description).Scan(&t.ID)
	if err != nil {
		log.Println("=== ADD TRANSACTION DB ERROR ===", err)
		http.Error(w, `{"error": "Failed to record transaction"}`, http.StatusInternalServerError)
		return
	}

	t.UserID = userID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func GetTransactionSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)

	var summary models.Summary
	query := `SELECT 
		COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0),
		COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0)
		FROM transactions WHERE user_id = $1`

	err := database.DB.QueryRow(query, userID).Scan(&summary.TotalIncome, &summary.TotalExpense)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch summary"}`, http.StatusInternalServerError)
		return
	}

	summary.NetProfit = summary.TotalIncome - summary.TotalExpense
	json.NewEncoder(w).Encode(summary)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)

	rows, err := database.DB.Query(`SELECT id, user_id, type, amount, COALESCE(description, ''), created_at FROM transactions WHERE user_id = $1 ORDER BY id DESC`, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch transactions"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	transactions := []models.Transaction{}
	for rows.Next() {
		var t models.Transaction
		rows.Scan(&t.ID, &t.UserID, &t.Type, &t.Amount, &t.Description, &t.CreatedAt)
		transactions = append(transactions, t)
	}

	json.NewEncoder(w).Encode(transactions)
}