package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	http.HandleFunc("/activities", GetActivities)

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Generar actividades dinámicas
func generateActivities(startDate string, numActivities, numDays int) map[string][]map[string]string {
	layout := "02-01-2006"
	start, _ := time.Parse(layout, startDate)

	activities := make(map[string][]map[string]string)

	for i := 0; i < numActivities; i++ {
		activityKey := fmt.Sprintf("activity_%d", i)

		var logs []map[string]string
		for j := 0; j < numDays; j++ {
			date := start.AddDate(0, 0, j).Format(layout)
			statuses := []string{"suck", "failed", "regular", "accomplished", "excellence"}
			status := statuses[j%len(statuses)] // Alternar estados

			logs = append(logs, map[string]string{
				"date":   date,
				"status": status,
			})
		}
		activities[activityKey] = logs
	}

	return activities
}

// Handler para devolver las actividades
func GetActivities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Para evitar problemas de CORS
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"activities": generateActivities("01-01-2025", 10, 30), // 10 actividades, 30 días
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error al generar JSON", http.StatusInternalServerError)
	}
}
