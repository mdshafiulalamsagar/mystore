package models

import "time"

// Task represents customer orders or pending shop tasks
type Task struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	CustomerName    string    `json:"customer_name"`
	TaskDescription string    `json:"task_description"`
	Status          string    `json:"status"` // PENDING or COMPLETED
	Deadline        string    `json:"deadline"`
	CreatedAt       time.Time `json:"created_at"`
}