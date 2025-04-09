package service

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	openapi "act-back/internal/openapi"
)

type ActivityService struct {
	DB *sql.DB
}

// enableCors agrega los headers CORS a la respuesta
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
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
		err := rows.Scan(&rawDate, &a.Actividad, &a.Status)
		if err != nil {
			log.Println("Error escaneando fila:", err)
			continue
		}
		a.Date = rawDate // ya está en formato yyyy-mm-dd
		activities = append(activities, a)
	}

	return activities, nil
}

// HandleGetActivitiesGrouped maneja la solicitud HTTP y devuelve las actividades agrupadas
func (s *ActivityService) HandleGetActivitiesGrouped(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
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
		grouped[activity.Actividad] = append(grouped[activity.Actividad], entry)
	}

	resp := openapi.GetActivitiesGroupedResponse{Activities: &grouped}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// HandleGetActivities maneja la solicitud HTTP y devuelve la lista plana de actividades
func (s *ActivityService) HandleGetActivities(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
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
