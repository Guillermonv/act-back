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

func (s *ActivityService) HandleGetActivities(w http.ResponseWriter, r *http.Request) {
    activities, err := s.GetActivities()
    if err != nil {
        http.Error(w, "Error al obtener actividades", http.StatusInternalServerError)
        return
    }

    resp := openapi.GetActivitiesResponse{Activities: activities}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}
