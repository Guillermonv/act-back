package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	openapi "act-back/internal/openapi"
)

type WeightService struct {
	DB *sql.DB
}

func (s *WeightService) GetWeights() (map[string][]openapi.WeightEntry, error) {
	rows, err := s.DB.Query("SELECT date, type, value FROM weights ORDER BY date ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	grouped := make(map[string][]openapi.WeightEntry)

	for rows.Next() {
		var date string
		var typ string
		var value float64 // directamente float64

		if err := rows.Scan(&date, &typ, &value); err != nil {
			log.Println("Error escaneando fila:", err)
			continue
		}

		entry := openapi.WeightEntry{
			Weight: value, // guarda directamente
			Date:   date,
		}

		switch typ {
		case "ideal":
			grouped["ideal"] = append(grouped["ideal"], entry)
		case "actual":
			grouped["current"] = append(grouped["current"], entry)
		}
	}

	return grouped, nil
}

// HandleGetWeightsList handles the HTTP request and returns weights grouped by type
func (s *WeightService) HandleGetWeightList(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}

	weights, err := s.GetWeights()
	if err != nil {
		http.Error(w, "Error al obtener pesos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := openapi.GetWeightsListResponse{
		Ideal:   weights["ideal"],
		Current: weights["current"],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *WeightService) HandleAddWeight(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == http.MethodOptions {
		return
	}

	var req openapi.AddWeightRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al parsear el cuerpo: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Ejecutar UPSERT
	_, err := s.DB.Exec(`
		INSERT INTO weights (date, type, value)
		VALUES (?, 'actual', ?)
		ON DUPLICATE KEY UPDATE value = VALUES(value)
	`, req.Date, req.Weight)

	if err != nil {
		http.Error(w, "Error al insertar o actualizar peso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := openapi.UpdateWeightResponse{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
