package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	ShopName  string    `json:"shop_name"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}