package services

import (
	"act-back/internal/openapi"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type TaskService struct {
	DB *sql.DB
}

// --- Middleware de CORS ---
func enableCORS(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400")

	// Si es preflight, respondemos directamente
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	}
	return false
}

// GET /tasks
func (s *TaskService) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	rows, err := s.DB.Query(`SELECT id, title, message, tag, status, start_date, end_date FROM tasks`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []openapi.Task
	for rows.Next() {
		var t openapi.Task
		err := rows.Scan(&t.Id, &t.Title, &t.Message, &t.Tag, &t.Status, &t.StartDate, &t.EndDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	writeJSON(w, tasks)
}

// GET /tasks/{id}
func (s *TaskService) HandleGetTaskByID(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	id := getIDFromPath(r.URL.Path, "/tasks/")
	if id == 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var t openapi.Task
	err := s.DB.QueryRow(`SELECT id, title, message, tag, status, start_date, end_date FROM tasks WHERE id = ?`, id).
		Scan(&t.Id, &t.Title, &t.Message, &t.Tag, &t.Status, &t.StartDate, &t.EndDate)
	if err == sql.ErrNoRows {
		http.Error(w, "Task no encontrada", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, t)
}

// POST /tasks
func (s *TaskService) HandleAddTask(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	var t openapi.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	_, err := s.DB.Exec(`INSERT INTO tasks (title, message, tag, status, start_date, end_date)
		VALUES (?, ?, ?, ?, ?, ?)`, t.Title, t.Message, t.Tag, t.Status, t.StartDate, t.EndDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{"message": "Task agregada correctamente"})
}

// PUT /tasks
func (s *TaskService) HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	var t openapi.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	_, err := s.DB.Exec(`UPDATE tasks SET title=?, message=?, tag=?, status=?, start_date=?, end_date=? WHERE id=?`,
		t.Title, t.Message, t.Tag, t.Status, t.StartDate, t.EndDate, t.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{"message": "Task actualizada correctamente"})
}

// DELETE /tasks
func (s *TaskService) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	if enableCORS(w, r) {
		return
	}

	id := getIDFromPath(r.URL.Path, "/tasks/")
	if id == 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	_, err := s.DB.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{"message": "Task eliminada correctamente"})
}

// Helpers
func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func getIDFromPath(path, prefix string) int64 {
	if len(path) <= len(prefix) {
		return 0
	}
	idStr := path[len(prefix):]
	var id int64
	_, err := fmt.Sscan(idStr, &id)
	if err != nil {
		return 0
	}
	return id
}
