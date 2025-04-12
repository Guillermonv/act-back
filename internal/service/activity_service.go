package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	openapi "act-back/internal/openapi"
)

type ActivityService struct {
	DB *sql.DB
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// GetActivities retrieves all activities from the database.
func (s *ActivityService) GetActivities() ([]openapi.Activity, error) {
	rows, err := s.DB.Query("SELECT fecha, actividad, status FROM actividades")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []openapi.Activity

	for rows.Next() {
		var a openapi.Activity
		var rawDate string
		err := rows.Scan(&rawDate, &a.Activity, &a.Status)
		if err != nil {
			log.Println("Error escaneando fila:", err)
			continue
		}
		// Formato dd-mm-yyyy
		parsedDate, err := time.Parse("2006-01-02", rawDate)
		if err != nil {
			a.Date = rawDate
		} else {
			a.Date = parsedDate.Format("02-01-2006")
		}
		activities = append(activities, a)
	}

	return activities, nil
}

// HandleGetActivitiesGrouped handles the HTTP request and returns grouped activities.
func (s *ActivityService) HandleGetActivitiesGrouped(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}

	activities, err := s.GetActivities()
	if err != nil {
		http.Error(w, "Error al obtener actividades: "+err.Error(), http.StatusInternalServerError)
		return
	}

	grouped := make(map[string][]openapi.ActivityDateStatus)

	for _, activity := range activities {
		entry := openapi.ActivityDateStatus{
			Date:   activity.Date,
			Status: activity.Status,
		}
		grouped[activity.Activity] = append(grouped[activity.Activity], entry)
	}

	resp := openapi.GetActivitiesGroupedResponse{Activities: &grouped}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// HandleGetActivities handles the HTTP request and returns flat activities.
func (s *ActivityService) HandleGetActivities(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}

	activities, err := s.GetActivities()
	if err != nil {
		http.Error(w, "Error al obtener actividades: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := openapi.GetActivitiesResponse{Activities: activities}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// HandleUpdateActivity handles the HTTP PUT request to update an activity's status.
func (s *ActivityService) HandleUpdateActivity(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var activity openapi.Activity
	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		http.Error(w, "Error de parseo JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// (Opcional) Debug: imprimir lo recibido
	fmt.Printf("Update: actividad=%s, fecha=%s, status=%s\n", activity.Activity, activity.Date, activity.Status)

	query := `
		UPDATE actividades
		SET status = ?
		WHERE fecha = ? AND actividad = ?`

	res, err := s.DB.Exec(query, activity.Status, activity.Date, activity.Activity)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar actividad: %v", err), http.StatusInternalServerError)
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "No se pudo determinar si se actualizó la actividad", http.StatusInternalServerError)
		return
	}

	if affected == 0 {
		http.Error(w, "No se encontró una actividad para actualizar", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Actividad actualizada correctamente"}`))
}
