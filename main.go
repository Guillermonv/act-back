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
	http.HandleFunc("/tasks", taskService.HandleGetTasks)
	http.HandleFunc("/tasks/add", taskService.HandleAddTask)
	http.HandleFunc("/tasks/update", taskService.HandleUpdateTask)
	http.HandleFunc("/tasks/delete/", taskService.HandleDeleteTask)
	http.HandleFunc("/tasks/get/", taskService.HandleGetTaskByID)

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
