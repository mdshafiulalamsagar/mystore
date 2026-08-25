package models

import "time"

type Order struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	CustomerName  string    `json:"customer_name"`
	CustomerPhone string    `json:"customer_phone"`
	TotalAmount   float64   `json:"total_amount"`
	Status        string    `json:"status"` // "Pending", "Delivered", "Cancelled"
	CreatedAt     time.Time `json:"created_at"`
}