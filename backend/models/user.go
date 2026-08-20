package models

import "time"

// User represents the store owner account
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	ShopName     string    `json:"shop_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}