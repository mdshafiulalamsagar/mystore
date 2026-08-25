package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"log"

	"golang.org/x/crypto/bcrypt"
	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
	"github.com/mdshafiulalamsagar/mystore/backend/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid input format"}`, http.StatusBadRequest)
		return
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" || user.Password == "" {
		http.Error(w, `{"error": "Email and password are required"}`, http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"error": "Failed to process password"}`, http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO users (name, email, password, shop_name) VALUES ($1, $2, $3, $4) RETURNING id`
	err = database.DB.QueryRow(query, user.Name, user.Email, string(hashedPassword), user.ShopName).Scan(&user.ID)
	if err != nil {
		log.Println("=== EXACT REGISTER DB ERROR ===", err)
		http.Error(w, `{"error": "Registration failed. Check server logs."}`, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error": "Invalid input format"}`, http.StatusBadRequest)
		return
	}

	input.Email = strings.TrimSpace(strings.ToLower(input.Email))

	var dbUser models.User
	var hashedPassword string
	query := `SELECT id, name, email, password, COALESCE(shop_name, '') FROM users WHERE LOWER(email) = LOWER($1)`
	err := database.DB.QueryRow(query, input.Email).Scan(&dbUser.ID, &dbUser.Name, &dbUser.Email, &hashedPassword, &dbUser.ShopName)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, `{"error": "Invalid email or password"}`, http.StatusUnauthorized)
			return
		}
		// Log exact database error in terminal
		log.Println("=== EXACT LOGIN DB ERROR ===", err)
		http.Error(w, `{"error": "Database error"}`, http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password))
	if err != nil {
		http.Error(w, `{"error": "Invalid email or password"}`, http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(dbUser.ID, dbUser.Email)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	dbUser.Password = ""
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"token":   token,
		"user":    dbUser,
	})
}