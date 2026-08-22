package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mdshafiulalamsagar/mystore/backend/database"
	"github.com/mdshafiulalamsagar/mystore/backend/middleware"
	"github.com/mdshafiulalamsagar/mystore/backend/models"
)

// CreateTask adds a new customer order or pending shop task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, `{"error": "Unauthorized access"}`, http.StatusUnauthorized)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, `{"error": "Invalid JSON input"}`, http.StatusBadRequest)
		return
	}

	query := `INSERT INTO tasks (user_id, customer_name, task_description, deadline) 
	          VALUES ($1, $2, $3, $4) RETURNING id, status, created_at`
	err := database.DB.QueryRow(query, userID, task.CustomerName, task.TaskDescription, task.Deadline).Scan(&task.ID, &task.Status, &task.CreatedAt)
	if err != nil {
		http.Error(w, `{"error": "Failed to create task"}`, http.StatusInternalServerError)
		return
	}

	task.UserID = userID
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Task created successfully",
		"task":    task,
	})
}

// GetTasks fetches all tasks for the logged in user
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, `{"error": "Unauthorized access"}`, http.StatusUnauthorized)
		return
	}

	query := `SELECT id, user_id, customer_name, task_description, status, deadline, created_at 
	          FROM tasks WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch tasks"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.UserID, &t.CustomerName, &t.TaskDescription, &t.Status, &t.Deadline, &t.CreatedAt); err != nil {
			http.Error(w, `{"error": "Error scanning task data"}`, http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, `{"error": "Error during task iteration"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}