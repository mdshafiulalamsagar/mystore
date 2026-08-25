package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)

	var o models.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	if o.Status == "" {
		o.Status = "Pending"
	}

	query := `INSERT INTO orders (user_id, customer_name, customer_phone, total_amount, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := database.DB.QueryRow(query, userID, o.CustomerName, o.CustomerPhone, o.TotalAmount, o.Status).Scan(&o.ID)
	if err != nil {
		http.Error(w, `{"error": "Failed to create order"}`, http.StatusInternalServerError)
		return
	}

	o.UserID = userID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)

	rows, err := database.DB.Query(`SELECT id, user_id, customer_name, COALESCE(customer_phone, ''), total_amount, status, created_at FROM orders WHERE user_id = $1 ORDER BY id DESC`, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch orders"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		var o models.Order
		rows.Scan(&o.ID, &o.UserID, &o.CustomerName, &o.CustomerPhone, &o.TotalAmount, &o.Status, &o.CreatedAt)
		orders = append(orders, o)
	}

	json.NewEncoder(w).Encode(orders)
}

func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := r.Context().Value(middleware.UserIDKey).(int)
	orderID := strings.TrimPrefix(r.URL.Path, "/orders/")

	var payload struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(`UPDATE orders SET status = $1 WHERE id = $2 AND user_id = $3`, payload.Status, orderID, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to update order status"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Order status updated"})
}