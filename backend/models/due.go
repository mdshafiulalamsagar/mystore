package models

import "time"

// Due represents shop dues like rent, bills, or vendor balance
type Due struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	DueDate   string    `json:"due_date"`
	Status    string    `json:"status"` // UNPAID or PAID
	CreatedAt time.Time `json:"created_at"`
}
