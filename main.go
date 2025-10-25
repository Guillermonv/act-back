package main

import (
	"act-back/internal/services"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "example:example@tcp(localhost:3306)/example")
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}

	activityService := &services.ActivityService{DB: db}
	weightService := &services.WeightService{DB: db}
	taskService := &services.TaskService{DB: db}

	// Endpoints existentes
	http.HandleFunc("/activities", activityService.HandleGetActivities)
	http.HandleFunc("/activities/grouped", activityService.HandleGetActivitiesGrouped)
	http.HandleFunc("/activities/update", activityService.HandleUpdateActivity)
	http.HandleFunc("/activities/populate", activityService.HandlePopulateActivities)

	http.HandleFunc("/weight/list", weightService.HandleGetWeightList)
	http.HandleFunc("/weight/add", weightService.HandleAddWeight)

	// Nuevos endpoints para Task
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskService.HandleGetTasks(w, r)
		case http.MethodPost:
			taskService.HandleAddTask(w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// /tasks/{id}
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskService.HandleGetTaskByID(w, r)
		case http.MethodPut:
			taskService.HandleUpdateTask(w, r)
		case http.MethodDelete:
			taskService.HandleDeleteTask(w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":80", nil))
}
