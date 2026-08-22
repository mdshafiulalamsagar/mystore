package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
)

// CreateTransaction adds a new income or expense record
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, `{"error": "Unauthorized access"}`, http.StatusUnauthorized)
		return
	}

	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, `{"error": "Invalid JSON input"}`, http.StatusBadRequest)
		return
	}

	// Validate transaction type
	if transaction.Type != "INCOME" && transaction.Type != "EXPENSE" {
		http.Error(w, `{"error": "Type must be either INCOME or EXPENSE"}`, http.StatusBadRequest)
		return
	}

	query := `INSERT INTO transactions (user_id, type, amount, category, description) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	err = database.DB.QueryRow(query, userID, transaction.Type, transaction.Amount, transaction.Category, transaction.Description).Scan(&transaction.ID, &transaction.CreatedAt)
	if err != nil {
		http.Error(w, `{"error": "Failed to record transaction"}`, http.StatusInternalServerError)
		return
	}

	transaction.UserID = userID
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "Transaction recorded successfully",
		"transaction": transaction,
	})
}

// GetTransactions retrieves all transactions for the authenticated user
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, `{"error": "Unauthorized access"}`, http.StatusUnauthorized)
		return
	}

	query := `SELECT id, user_id, type, amount, category, description, created_at 
	          FROM transactions WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch transactions"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	transactions := []models.Transaction{}
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.ID, &t.UserID, &t.Type, &t.Amount, &t.Category, &t.Description, &t.CreatedAt)
		if err != nil {
			http.Error(w, `{"error": "Error scanning transaction row"}`, http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, t)
	}

	json.NewEncoder(w).Encode(transactions)
}

// GetFinancialSummary calculates total income, total expenses, and net profit
func GetFinancialSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, `{"error": "Unauthorized access"}`, http.StatusUnauthorized)
		return
	}

	var totalIncome, totalExpense float64

	// Fetch sum of income
	incomeQuery := `SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE user_id = $1 AND type = 'INCOME'`
	database.DB.QueryRow(incomeQuery, userID).Scan(&totalIncome)

	// Fetch sum of expenses
	expenseQuery := `SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE user_id = $1 AND type = 'EXPENSE'`
	database.DB.QueryRow(expenseQuery, userID).Scan(&totalExpense)

	netProfit := totalIncome - totalExpense

	json.NewEncoder(w).Encode(map[string]interface{}{
		"total_income":  totalIncome,
		"total_expense": totalExpense,
		"net_profit":    netProfit,
	})
}