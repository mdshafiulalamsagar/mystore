package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
	"github.com/mdshafiulalamsagar/mystore/backend/utils"
)

// RegisterUser handles creation of a new store owner account
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, `{"error": "Invalid JSON input"}`, http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		http.Error(w, `{"error": "Error hashing password"}`, http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO users (name, shop_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id`
	err = database.DB.QueryRow(query, user.Name, user.ShopName, user.Email, hashedPassword).Scan(&user.ID)
	if err != nil {
		http.Error(w, `{"error": "User registration failed or email exists"}`, http.StatusBadRequest)
		return
	}

	user.PasswordHash = ""
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user":    user,
	})
}

// LoginUser authenticates a user and returns a JWT token
func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	var user models.User
	query := `SELECT id, name, shop_name, email, password_hash FROM users WHERE email = $1`
	row := database.DB.QueryRow(query, input.Email)
	err = row.Scan(&user.ID, &user.Name, &user.ShopName, &user.Email, &user.PasswordHash)

	if err == sql.ErrNoRows || !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		http.Error(w, `{"error": "Invalid email or password"}`, http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}