package models

import "time"

type Due struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	CustomerName string    `json:"customer_name"`
	Phone        string    `json:"phone"`
	TotalDue     float64   `json:"total_due"`
	PaidAmount   float64   `json:"paid_amount"`
	Status       string    `json:"status"` // "Unpaid", "Partial", "Paid"
	CreatedAt    time.Time `json:"created_at"`
}